package auth

import (
	"errors"
	"time"

	"elkhawaga-erp/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	BranchID *uint  `json:"branch_id,omitempty"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func GenerateTokenPair(userID uint, username string, branchID *uint) (*TokenPair, error) {
	cfg := config.Get()

	// إنشاء Access Token
	accessClaims := &Claims{
		UserID:   userID,
		Username: username,
		BranchID: branchID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.JWT.Expires)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "elkhawaga-erp",
			Subject:   username,
			ID:        uuid.New().String(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	// إنشاء Refresh Token
	refreshToken := uuid.New().String()

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(cfg.JWT.Expires.Seconds()),
	}, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.Get().JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func RefreshToken(refreshToken string) (*TokenPair, error) {
	// التحقق من صحة Refresh Token في قاعدة البيانات
	// هذا سيتم تنفيذه في خدمة الجلسات
	return nil, errors.New("not implemented")
}
