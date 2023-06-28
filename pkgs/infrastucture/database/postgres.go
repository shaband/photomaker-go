package database

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type PostgresConnnection struct {
	Host     string
	User     string
	Password string
	DB_NAME  string
	Port     string
	SslMode  string
	TimeZone string
}

func (connection PostgresConnnection) Connect(config *gorm.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(`
	host=%v 
	user=%v 
	password=%v 
	dbname=%v 
	port=%v 
	sslmode=%v 
	TimeZone=%v`,
	 connection.Host, connection.User, connection.Password, connection.DB_NAME, connection.Port, connection.SslMode, connection.TimeZone)
	return gorm.Open(postgres.Open(dsn), config)
}

func NewPostgresConnection() PostgresConnnection {

	return PostgresConnnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DB_NAME: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SslMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
}
