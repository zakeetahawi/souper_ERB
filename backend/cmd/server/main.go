package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"elkhawaga-erp/internal/auth"
	"elkhawaga-erp/internal/config"
	"elkhawaga-erp/internal/db"
	"elkhawaga-erp/internal/modules"
	"elkhawaga-erp/internal/modules/customers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// تحميل الإعدادات
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// الاتصال بقاعدة البيانات
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// إنشاء تطبيق Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	// إضافة Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Device-Fingerprint",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// تسجيل الموديولات
	moduleRegistry := modules.NewModuleRegistry()
	if err := moduleRegistry.InitializeDefaultModules(); err != nil {
		log.Fatal("Failed to initialize modules:", err)
	}

	// تسجيل مسارات API
	registerAPIRoutes(app, moduleRegistry)

	// مسار الصحة
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"app":    config.Get().App.Name,
		})
	})

	// تشغيل الخادم
	cfg := config.Get()
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

	log.Printf("Server started on port %d", cfg.App.Port)

	// انتظار إشارة الإغلاق
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatal("Failed to shutdown server:", err)
	}
}

// registerAPIRoutes - تسجيل مسارات API
func registerAPIRoutes(app *fiber.App, moduleRegistry *modules.ModuleRegistry) {
	// مسارات المصادقة
	authAPI := auth.NewAuthAPI()
	authAPI.RegisterRoutes(app)

	// مسارات العملاء
	if moduleRegistry.IsModuleEnabled("customers") {
		customerAPI := customers.NewCustomerAPI()
		customerAPI.RegisterRoutes(app)
	}

	// مسارات الموديولات الأخرى (قيد التطوير)
	registerPlaceholderRoutes(app)
}

// registerPlaceholderRoutes - تسجيل مسارات placeholder للموديولات قيد التطوير
func registerPlaceholderRoutes(app *fiber.App) {
	modules := []string{
		"orders", "inventory", "field-survey", "factory", "installations",
		"data-sync", "maintenance", "development", "display-tuning",
		"sales", "marketing", "projects",
	}

	for _, module := range modules {
		app.Get(fmt.Sprintf("/api/%s", module), func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"success": false,
				"error":   "Module is under development",
				"message": "هذا الموديول قيد التطوير",
			})
		})
	}
}

// customErrorHandler - معالج الأخطاء المخصص
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}
