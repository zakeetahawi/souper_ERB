package auth

import (
	"elkhawaga-erp/internal/db"
	"elkhawaga-erp/internal/db/repo"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthAPI struct {
	db             *gorm.DB
	sessionService *SessionService
	rbacService    *RBACService
}

func NewAuthAPI() *AuthAPI {
	return &AuthAPI{
		db:             db.GetDB(),
		sessionService: NewSessionService(),
		rbacService:    NewRBACService(),
	}
}

// RegisterRoutes - تسجيل مسارات المصادقة
func (api *AuthAPI) RegisterRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")

	auth.Post("/login", api.Login)
	auth.Post("/refresh", api.RefreshToken)
	auth.Post("/logout", api.Logout)
	auth.Get("/me", api.GetCurrentUser)
}

// LoginRequest - طلب تسجيل الدخول
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse - استجابة تسجيل الدخول
type LoginResponse struct {
	Success bool `json:"success"`
	Data    struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64  `json:"expires_in"`
		User         struct {
			ID       uint     `json:"id"`
			Username string   `json:"username"`
			BranchID *uint    `json:"branch_id,omitempty"`
			Roles    []string `json:"roles"`
		} `json:"user"`
	} `json:"data"`
}

// Login - تسجيل الدخول
func (api *AuthAPI) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid request body",
		})
	}

	// التحقق من وجود المستخدم
	var user repo.User
	err := api.db.Preload("Roles").Where("username = ? AND is_active = ?", req.Username, true).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "invalid credentials",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "database error",
		})
	}

	// التحقق من كلمة المرور
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "invalid credentials",
		})
	}

	// إنشاء التوكن
	tokenPair, err := GenerateTokenPair(user.ID, user.Username, user.BranchID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to generate token",
		})
	}

	// إنشاء الجلسة
	deviceFingerprint := c.Get("X-Device-Fingerprint", "")
	ip := c.IP()
	userAgent := c.Get("User-Agent", "")

	if err := api.sessionService.CreateSession(user.ID, deviceFingerprint, tokenPair.RefreshToken, ip, userAgent); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to create session",
		})
	}

	// تحضير الأدوار
	roles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.Name
	}

	// إعداد الاستجابة
	response := LoginResponse{
		Success: true,
	}
	response.Data.AccessToken = tokenPair.AccessToken
	response.Data.RefreshToken = tokenPair.RefreshToken
	response.Data.ExpiresIn = tokenPair.ExpiresIn
	response.Data.User.ID = user.ID
	response.Data.User.Username = user.Username
	response.Data.User.BranchID = user.BranchID
	response.Data.User.Roles = roles

	return c.JSON(response)
}

// RefreshTokenRequest - طلب تجديد التوكن
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// RefreshToken - تجديد التوكن
func (api *AuthAPI) RefreshToken(c *fiber.Ctx) error {
	var req RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid request body",
		})
	}

	// التحقق من صحة Refresh Token
	claims, err := ValidateToken(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "invalid refresh token",
		})
	}

	// التحقق من وجود الجلسة
	session, err := api.sessionService.ValidateSession(claims.UserID, req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "session not found or expired",
		})
	}

	// إنشاء توكن جديد
	tokenPair, err := GenerateTokenPair(claims.UserID, claims.Username, claims.BranchID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to generate token",
		})
	}

	// تحديث الجلسة
	if err := api.sessionService.RevokeSession(session.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to update session",
		})
	}

	deviceFingerprint := c.Get("X-Device-Fingerprint", "")
	ip := c.IP()
	userAgent := c.Get("User-Agent", "")

	if err := api.sessionService.CreateSession(claims.UserID, deviceFingerprint, tokenPair.RefreshToken, ip, userAgent); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to create new session",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"access_token":  tokenPair.AccessToken,
			"refresh_token": tokenPair.RefreshToken,
			"expires_in":    tokenPair.ExpiresIn,
		},
	})
}

// Logout - تسجيل الخروج
func (api *AuthAPI) Logout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	deviceFingerprint := c.Get("X-Device-Fingerprint", "")

	// إلغاء الجلسة
	if deviceFingerprint != "" {
		if err := api.sessionService.RevokeSessionByDevice(userID, deviceFingerprint); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   "failed to logout",
			})
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"message": "تم تسجيل الخروج بنجاح",
		},
	})
}

// GetCurrentUser - جلب بيانات المستخدم الحالي
func (api *AuthAPI) GetCurrentUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user repo.User
	err := api.db.Preload("Roles").Preload("Branch").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "user not found",
		})
	}

	// تحضير الأدوار
	roles := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = role.Name
	}

	// جلب الموديولات المتاحة
	modules, err := api.rbacService.GetUserModules(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to get user modules",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"branch": fiber.Map{
				"id":   user.BranchID,
				"name": user.Branch.NameAr,
			},
			"roles":   roles,
			"modules": modules,
		},
	})
}

// AuthMiddleware - Middleware للمصادقة
func (api *AuthAPI) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "authorization header required",
			})
		}

		// استخراج التوكن
		token := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			token = authHeader[7:]
		}

		// التحقق من صحة التوكن
		claims, err := ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "invalid token",
			})
		}

		// إضافة بيانات المستخدم للسياق
		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)
		c.Locals("branch_id", claims.BranchID)

		return c.Next()
	}
}

// RBACMiddleware - Middleware للصلاحيات
func (api *AuthAPI) RBACMiddleware(moduleKey, action string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)

		hasPermission, err := api.rbacService.CheckPermission(userID, moduleKey, action)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   "failed to check permissions",
			})
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"error":   "insufficient permissions",
			})
		}

		return c.Next()
	}
}

// CreateUser - إنشاء مستخدم جديد (للتطوير فقط)
func (api *AuthAPI) CreateUser(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		BranchID *uint  `json:"branch_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "invalid request body",
		})
	}

	// تشفير كلمة المرور
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to hash password",
		})
	}

	user := repo.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		BranchID:     req.BranchID,
		IsActive:     true,
	}

	if err := api.db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"message":  "تم إنشاء المستخدم بنجاح",
		},
	})
}
