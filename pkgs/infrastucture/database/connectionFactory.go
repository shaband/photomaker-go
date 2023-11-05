package database

import (
	"errors"
	"gorm.io/gorm"
)

type ConnectionFactory interface {
	Connect(config *gorm.Config) (*gorm.DB, error)
}

func connectionFactory(factory string) (ConnectionFactory, error) {
	if factory == "postgres" {
		return NewPostgresConnection(), nil
	} else {
		return nil, errors.New("Invalid Connection")
	}
}
