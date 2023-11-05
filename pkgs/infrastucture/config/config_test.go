package config

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Load environment variables from test .env file
	godotenv.Load(".env.test")

	// Call LoadConfig function
	config, err := LoadConfig()

	// Check that LoadConfig returns the correct Config struct and no error
	assert.NoError(t, err)
	assert.Equal(t, "test_db_name", config.DSN)

	// Unload environment variables and check that LoadConfig returns an error
	godotenv.Overload(".env.empty")
	_, err = LoadConfig()
	assert.Error(t, err)
}
