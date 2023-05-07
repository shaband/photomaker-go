package contacts

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Name         string
	Phone        string
	Email        string
	File         string
	Category     string
	ServiceTypes interface{} `gorm:"serializer:json"`
}
