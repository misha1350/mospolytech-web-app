package config

import (
	"os"
	"strconv"
	"time"
)

type DatabaseConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func LoadDatabaseConfig() *DatabaseConfig {
	maxOpen, _ := strconv.Atoi(getEnvOrDefault("DB_MAX_OPEN_CONNS", "25"))
	maxIdle, _ := strconv.Atoi(getEnvOrDefault("DB_MAX_IDLE_CONNS", "5"))
	maxLifetime, _ := strconv.Atoi(getEnvOrDefault("DB_CONN_MAX_LIFETIME_MINUTES", "5"))

	return &DatabaseConfig{
		MaxOpenConns:    maxOpen,
		MaxIdleConns:    maxIdle,
		ConnMaxLifetime: time.Duration(maxLifetime) * time.Minute,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
