package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"index;unique"`
	Email    string `gorm:"index;unique"`
	Password string
	Phone    string
	Image    string
}
