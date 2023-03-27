package contacts

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name         string
	Category     string
	Phone        string
	Email        string
	File         string
	ServiceTypes interface{} `gorm:"serializer:json"`
}
