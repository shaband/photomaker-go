package categories

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

type CategoryHandler struct {
	service    *CategoriesService
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler CategoryHandler) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin.Categories.index.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
		"Categories": handler.service.GetAll(),
	}))

}

// func (handler CategoryHandler) Show(c *gin.Context)   {

// }
func (handler CategoryHandler) Create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.Categories.create.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
	}))
}
func (handler CategoryHandler) Store(ctx *gin.Context) {
	CategoryRequest := CategoryRequest{}
	ctx.ShouldBind(&CategoryRequest)
	errs := validator.Validate(CategoryRequest)
	if errs != nil {
		helpers.RedirectFailedWithValidation(ctx, "/admin/Categories/create", errs, CategoryRequest)
		return
	}
	handler.service.Store(&CategoryRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/Categories", "Category created successfully")

}
func (handler CategoryHandler) Edit(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/Categories", "Invalid User ")
		return
	}
	ctx.HTML(http.StatusOK, "admin.users.edit.gohtml", gin.H{
		"token": csrf.GetToken(ctx),
		"user":  handler.service.Find(id),
	})
}

func (handler UserHandler) Update(ctx *gin.Context) {
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
func (handler UserHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/users", "Invalid User ")
		return
	}
	handler.service.DeleteById(id)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/users", "Deleted Successfully")

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
