package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/gorm"
)

type UserHandler struct {
	service    *UserService
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler UserHandler) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin.users.index.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
		"users": handler.service.GetAll(),
	}))

}

// func (handler UserHandler) Show(c *gin.Context)   {

// }
func (handler UserHandler) Create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.users.create.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
	}))
}
func (handler UserHandler) Store(ctx *gin.Context) {
	UserRequest := UserRequest{}
	ctx.ShouldBind(&UserRequest)
	errs := validator.Validate(UserRequest)
	if errs != nil {
		helpers.RedirectFailedWithValidation(ctx, "/admin/users/create", errs, UserRequest)
		return
	}
	handler.service.Store(&UserRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/users", "User created successfully")

}
func (handler UserHandler) Edit(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("user"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "admin/users", "Invalid User ")
		return
	}
	ctx.HTML(http.StatusOK, "admin.users.edit.gohtml", gin.H{
		"token": csrf.GetToken(ctx),
		"user":  handler.service.Find(id),
	})
}

func (handler UserHandler) Update(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("user"))
	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/users/%s/edit", ctx.Param("user")), "Invalid User ")
		return 
	}
	if err != nil {
		panic(err)
	}
	UserRequest := UserRequest{}
	ctx.ShouldBind(&UserRequest)

	errs := validator.Validate(UserRequest)
	if errs != nil {
		helpers.RedirectFailedWithValidation(ctx, fmt.Sprintf("/admin/users/%s/edit", ctx.Param("user")), errs, UserRequest)
		return
	}
	handler.service.Update(id, UserRequest.ToEntity())
	helpers.RedirectSuccessWithMessage(ctx, "/admin/users", "User Updated successfully")

	helpers.RedirectSuccessWithMessage(ctx, "admin/users", "Updated Successfully")

}
func (handler UserHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("user"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "admin/users", "Invalid User ")
		return
	}
	handler.service.DeleteById(id)
	helpers.RedirectSuccessWithMessage(ctx, "/users/index", "Deleted Successfully")

}

func NewUserHandler(DB *gorm.DB, commonData func(c *gin.Context, data gin.H) gin.H) *UserHandler {

	return &UserHandler{
		service:    NewUserService(DB),
		commonData: commonData,
	}
}

func withCommonData(c *gin.Context, data gin.H) gin.H {

	commonData := c.MustGet("CommonData").(gin.H)
	data["commonData"] = commonData

	return data
}
