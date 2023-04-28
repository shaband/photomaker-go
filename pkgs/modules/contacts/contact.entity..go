package contacts

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Name         string               `form:"name"`
	Phone        string               `form:"phone"`
	Email        string               `form:"email"`
	File         multipart.FileHeader `from:"attachment"`
	Category     string               `form:"subject"`
	ServiceTypes interface{}          `gorm:"serializer:json" form:"service_types[]"`
}
