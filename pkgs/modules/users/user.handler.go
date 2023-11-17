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

type Handler struct {
	service    *UserService
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler Handler) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin.users.index.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
		"users": handler.service.GetAll(),
	}))

}

// func (handler Handler) Show(c *gin.Context)   {

// }
func (handler Handler) Create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.users.create.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
	}))
}
func (handler Handler) Store(ctx *gin.Context) {
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
func (handler Handler) Edit(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/users", "Invalid User ")
		return
	}
	ctx.HTML(http.StatusOK, "admin.users.edit.gohtml", gin.H{
		"token": csrf.GetToken(ctx),
		"user":  handler.service.Find(id),
	})
}

func (handler Handler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	fmt.Println(id)
	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/users/%s/edit", ctx.Param("id")), "Invalid User ")
		return
	}
	UserRequest := UserRequest{}
	ctx.ShouldBind(&UserRequest)

	errs := validator.Validate(UserRequest)
	if errs != nil {
		fmt.Println(errs)

		helpers.RedirectFailedWithValidation(ctx, fmt.Sprintf("/admin/users/%s/edit", ctx.Param("id")), errs, UserRequest)
		return
	}
	handler.service.Update(uint(id), UserRequest.ToEntity())
	helpers.RedirectSuccessWithMessage(ctx, "/admin/users", "User Updated successfully")

}
func (handler Handler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/users", "Invalid User ")
		return
	}
	handler.service.DeleteById(id)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/users", "Deleted Successfully")

}

func NewHandler(DB *gorm.DB, commonData func(c *gin.Context, data gin.H) gin.H) *Handler {

	return &Handler{
		service:    NewUserService(DB),
		commonData: commonData,
	}
}

func withCommonData(c *gin.Context, data gin.H) gin.H {

	commonData := c.MustGet("CommonData").(gin.H)
	data["commonData"] = commonData

	return data
}
