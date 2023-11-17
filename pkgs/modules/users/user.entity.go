package users

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"index;unique"`
	Email    string `gorm:"index;unique"`
	Password string
	Phone    string
	Image    string
}

func (u *User) CheckPassword(password string) bool {
	return CheckPasswordHash(u.Password, password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
