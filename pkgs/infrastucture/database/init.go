package database

import (
	"os"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type ConfigLoader interface {
	GetEnv(key string) string
}

func Init(configLoader ConfigLoader) {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	connection, err := connectionFactory(os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic(err)
	}
	DB, err := connection.Connect(config)
	if err != nil {
		panic("failed to connect database")
	}
	db = DB
}
