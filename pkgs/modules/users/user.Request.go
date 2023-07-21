package users

type UserRequest struct {
	Username string `form:"username" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
	Phone    string `form:"password" validate:"required"`
}
