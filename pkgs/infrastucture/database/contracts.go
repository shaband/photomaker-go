package database

import "gorm.io/gorm"

type connection interface {
	Connect(config *gorm.Config) (*gorm.DB, error)
}

