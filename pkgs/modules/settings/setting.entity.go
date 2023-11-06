package settings

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Slug  string `gorm:"index;unique"`
	Label string
	Value string
	Page  string
	Type  string
}
