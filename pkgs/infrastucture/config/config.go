package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
	EnvLoader EnvLoader
}
func (c *Config) GetEnv(key string) string {

	return c.EnvLoader.GetEnv(key)

}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	dbname := os.Getenv("DB_NAME")
	// pass := os.Getenv("DB_PASS")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbname := os.Getenv("DB_NAME")

	return &Config{DSN: dbname}, nil
}

type EnvLoader interface {
	GetEnv(key string) string
}