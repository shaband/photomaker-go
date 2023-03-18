package settings

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	Label string
	Value string
	Page  string
	Type  string
}
