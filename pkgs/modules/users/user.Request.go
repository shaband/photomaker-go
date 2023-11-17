package users

type UserRequest struct {
	Username string `form:"username" validate:"required,alphanum"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
	Phone    string `form:"phone" validate:"required"`
}

func (u *UserRequest) ToEntity() *User {

	hashedPassword, _ := HashPassword(u.Password)
	return &User{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
		Phone:    u.Phone,
	}

}
