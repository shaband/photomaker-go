package users

type UserRequest struct {
	ID int	`gorm:"index;unique"`
	Username string `gorm:"index;unique"`
	Email    string `gorm:"index;unique"`
	Password string
	Phone    string 
	Image    string ``

}
