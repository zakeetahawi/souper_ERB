package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	DB       DBConfig
	Redis    RedisConfig
	JWT      JWTConfig
	Features FeaturesConfig
	Files    FilesConfig
}

type AppConfig struct {
	Name    string
	Env     string
	Port    int
	BaseURL string
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

type RedisConfig struct {
	Host string
	Port int
	DB   int
}

type JWTConfig struct {
	Secret         string
	Expires        time.Duration
	RefreshExpires time.Duration
}

type FeaturesConfig struct {
	ImportExport            bool
	GoogleSheetsSync        bool
	SingleDeviceLogin       bool
	DeviceFingerprintHeader string
}

type FilesConfig struct {
	Storage         string
	UploadsDir      string
	CompanyLogoPath string
	HeaderLogoPath  string
}

var GlobalConfig *Config

func Load() error {
	if err := godotenv.Load(); err != nil {
		// لا نعتبر هذا خطأ إذا لم يكن ملف .env موجود
	}

	GlobalConfig = &Config{
		App: AppConfig{
			Name:    getEnv("APP_NAME", "zakeeERP"),
			Env:     getEnv("APP_ENV", "development"),
			Port:    getEnvAsInt("APP_PORT", 8080),
			BaseURL: getEnv("APP_BASE_URL", "http://localhost:8080"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			Name:     getEnv("DB_NAME", "zakee_erp"),
			User:     getEnv("DB_USER", "erp_user"),
			Password: getEnv("DB_PASS", "password"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Redis: RedisConfig{
			Host: getEnv("REDIS_HOST", "localhost"),
			Port: getEnvAsInt("REDIS_PORT", 6379),
			DB:   getEnvAsInt("REDIS_DB", 0),
		},
		JWT: JWTConfig{
			Secret:         getEnv("JWT_SECRET", "change_me_in_production"),
			Expires:        getEnvAsDuration("JWT_EXPIRES", 15*time.Minute),
			RefreshExpires: getEnvAsDuration("REFRESH_EXPIRES", 720*time.Hour),
		},
		Features: FeaturesConfig{
			ImportExport:            getEnvAsBool("FEATURE_IMPORT_EXPORT", true),
			GoogleSheetsSync:        getEnvAsBool("FEATURE_GOOGLE_SHEETS_SYNC", true),
			SingleDeviceLogin:       getEnvAsBool("SINGLE_DEVICE_LOGIN", true),
			DeviceFingerprintHeader: getEnv("DEVICE_FINGERPRINT_HEADER", "X-Device-Fingerprint"),
		},
		Files: FilesConfig{
			Storage:         getEnv("FILES_STORAGE", "local"),
			UploadsDir:      getEnv("UPLOADS_DIR", "/data/uploads"),
			CompanyLogoPath: getEnv("COMPANY_LOGO_PATH", "/data/uploads/company/logo.png"),
			HeaderLogoPath:  getEnv("HEADER_LOGO_PATH", "/data/uploads/company/header_logo.png"),
		},
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func Get() *Config {
	return GlobalConfig
}
