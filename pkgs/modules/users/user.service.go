package users

import (
	// "github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (service *UserService) GetAll() []User {
	Users := []User{}
	service.db.Find(&Users)
	return Users
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
func (service UserService) Update(ID int, user *User) {

	service.db.Save(user)
}

func (service *UserService) DeleteById(ID int) *User {
	// user:=service.Find(ID)
	service.db.Delete(&User{}, ID)
	return nil
}

// func (service *UserService) ValidateAndStore(userRequest *UserRequest) (*User, *map[string]string) {

// 	errs := validator.Validate(userRequest)
// 	if errs != nil {

// 		return nil, &errs
// 	}
// 	user := userRequest.ToEntity()
// 	service.db.Create(user)

// 	return user, nil
// }

func NewUserService(db *gorm.DB) *UserService {

	return &UserService{
		db: db,
	}
}
