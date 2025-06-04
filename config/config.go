package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Config 結構體包含所有應用程序配置
type Config struct {
	Server    ServerConfig
	MongoDB   MongoDBConfig
	JWT       JWTConfig
	Logging   LoggingConfig
	RateLimit RateLimitConfig
	CORS      CORSConfig
}

// ServerConfig 包含服務器相關配置
type ServerConfig struct {
	Port    string
	GinMode string
}

// MongoDBConfig 包含 MongoDB 相關配置
type MongoDBConfig struct {
	URI      string
	Database string
	Username string
	Password string
	Timeout  time.Duration
}

// JWTConfig 包含 JWT 相關配置
type JWTConfig struct {
	SecretKey       string
	ExpirationHours int
}

// LoggingConfig 包含日誌相關配置
type LoggingConfig struct {
	Level string
	File  string
}

// RateLimitConfig 包含 API 限流相關配置
type RateLimitConfig struct {
	Requests int
	Window   int
}

// CORSConfig 包含 CORS 相關配置
type CORSConfig struct {
	AllowedOrigins []string
}

// LoadConfig 從環境變數加載配置
func LoadConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		MongoDB: MongoDBConfig{
			URI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
			Database: getEnv("MONGODB_DATABASE", "go_api_db"),
			Username: getEnv("MONGODB_USERNAME", ""),
			Password: getEnv("MONGODB_PASSWORD", ""),
			Timeout:  time.Duration(getEnvAsInt("MONGODB_TIMEOUT", 10)) * time.Second,
		},
		JWT: JWTConfig{
			SecretKey:       getEnv("JWT_SECRET_KEY", "default_secret_key"),
			ExpirationHours: getEnvAsInt("JWT_EXPIRATION_HOURS", 24),
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "debug"),
			File:  getEnv("LOG_FILE", "./logs/app.log"),
		},
		RateLimit: RateLimitConfig{
			Requests: getEnvAsInt("RATE_LIMIT_REQUESTS", 100),
			Window:   getEnvAsInt("RATE_LIMIT_WINDOW", 1),
		},
		CORS: CORSConfig{
			AllowedOrigins: getEnvAsStringSlice("ALLOWED_ORIGINS", []string{"http://localhost:3000", "http://localhost:8080"}),
		},
	}
}

// getEnv 獲取環境變數，如果不存在則返回默認值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt 獲取整數類型的環境變數
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsStringSlice 獲取字符串切片類型的環境變數
func getEnvAsStringSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}
