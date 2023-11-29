package users

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (service *UserService) GetAll() *[]User {
	Users := []User{}
	service.db.Find(&Users)
	return &Users
}
func (service *UserService) FindUserByEmail(Email string) *User {
	User := User{}
	service.db.Where("email = ?", Email).Find(&User)
	return &User
}

func (service *UserService) Find(ID int) *User {
	User := User{}
	service.db.Find(&User, ID)
	return &User
}
func (service UserService) Update(ID uint, user *User) {
	user.ID = ID
	result := service.db.Save(user)
	fmt.Println(result)
}

func (service *UserService) DeleteById(ID int) *User {
	// user:=service.Find(ID)
	service.db.Delete(&User{}, ID)
	return nil
}

func (service *UserService) Store(userRequest *UserRequest) *User {

	user := userRequest.ToEntity()
	service.db.Create(user)

	return user
}

func (service *UserService) Login(ctx *gin.Context, LoginRequest loginRequest) error {

	user := service.FindUserByEmail(LoginRequest.Email)
	fmt.Println(LoginRequest)
	if valid := CheckPasswordHash(LoginRequest.Password, user.Password); !valid {

		return errors.New("wrong Password")
	}
	s := sessions.Default(ctx)

	s.Set("admin", strconv.FormatUint(uint64(user.ID), 10))
	return s.Save()
}

func NewUserService(db *gorm.DB) *UserService {

	return &UserService{
		db: db,
	}
}
