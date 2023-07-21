package users

import (
	"net/http"
	"strconv"

	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *UserService
}

func (handler UserHandler) Index(c *gin.Context) {

	c.HTML(http.StatusOK, "admin.users.index.gohtml", gin.H{
		"users": handler.service.GetAll(),
	})

}

// func (handler UserHandler) Show(c *gin.Context)   {

// }
func (handler UserHandler) Create(c *gin.Context) {

	c.HTML(http.StatusOK, "admin.users.create.gohtml", gin.H{})
}
func (handler UserHandler) Store(c *gin.Context) {

	helpers.RedirectSuccessWithMessage(c, "/users/index", "Saved Successfully")
}
func (handler UserHandler) Edit(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("user"))
	if err != nil {

		panic(err)
	}
	c.HTML(http.StatusOK, "admin.users.edit.gohtml", gin.H{
		"users": handler.service.Find(id),
	})
}

func (handler UserHandler) Update(c *gin.Context) {

	helpers.RedirectSuccessWithMessage(c, "/users/index", "Updated Successfully")

}
func (handler UserHandler) Delete(c *gin.Context) {

	helpers.RedirectSuccessWithMessage(c, "/users/index", "Deleted Successfully")

}

func NewUserHandler(DB *gorm.DB) *UserHandler {

	return &UserHandler{
		service: NewUserService(DB),
	}
}
