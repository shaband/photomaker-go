package database

import "gorm.io/gorm"

func GetConnection() *gorm.DB {
	return db
}
