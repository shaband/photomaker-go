package users

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/gorm"
)

type AuthHandler struct {
	service *UserService
	db      *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		db:      db,
		service: NewUserService(db),
	}

}
func (handler *AuthHandler) GetLoginPage(c *gin.Context) {

	c.HTML(http.StatusOK, "admin.auth.login.gohtml",
		gin.H{
			"token": csrf.GetToken(c),
		})
}

// login is a handler that parses a form and checks for specific data.
func (handler *AuthHandler) Login(ctx *gin.Context) {
	var LoginRequest loginRequest
	ctx.ShouldBind(&LoginRequest)
	if errors := validator.Validate(LoginRequest); errors != nil {
		helpers.RedirectFailedWithValidation(ctx, "/admin/auth/login", errors, LoginRequest)
		return
	}
	err := handler.service.Login(ctx, LoginRequest)
	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, "/admin/auth/login", err.Error())
		return
	}

	helpers.RedirectSuccessWithMessage(ctx, "/admin/users", "Successfully logged in")

}

// logout is the handler called for the user to log out.
func (handler *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("admin")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete("admin")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

// me is the handler that will return the user information stored in the
// session.
func (handler *AuthHandler) Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("admin")
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// status is the handler that will tell the user whether it is logged in or not.
func (handler *AuthHandler) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

type loginRequest struct {
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}
