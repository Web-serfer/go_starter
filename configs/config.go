package config

import (
	"os"
)

// Config структура для хранения конфигурации приложения
type Config struct {
	ServerPort string
	DatabaseURL string
	Debug bool
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() *Config {
	return &Config{
		ServerPort: getEnvOrDefault("SERVER_PORT", ":8080"),
		DatabaseURL: getEnvOrDefault("DATABASE_URL", "sqlite://local.db"),
		Debug: getEnvOrDefault("DEBUG", "false") == "true",
	}
}

// getEnvOrDefault возвращает значение переменной окружения или значение по умолчанию
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}