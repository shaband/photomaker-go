package categories

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	csrf "github.com/utrack/gin-csrf"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	service    *Service
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler CategoryHandler) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin.categories.index.gohtml", withCommonData(ctx, gin.H{
		"token":      csrf.GetToken(ctx),
		"categories": handler.service.GetAll(),
	}))

}

// func (handler CategoryHandler) Show(c *gin.Context)   {

// }
func (handler CategoryHandler) Create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.categories.create.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
	}))
}
func (handler CategoryHandler) Store(ctx *gin.Context) {
	CategoryRequest := CreateCategoryRequest{}
	_ = ctx.ShouldBind(&CategoryRequest)

	errs := validator.Validate(CategoryRequest)
	if errs != nil {
		helpers.RedirectFailedWithValidation(ctx, "/admin/categories/create", errs, CategoryRequest)
		return
	}

	handler.service.Store(ctx, &CategoryRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/categories", "Category created successfully")

}
func (handler CategoryHandler) Edit(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/categories", "Invalid Category ")
		return
	}
	ctx.HTML(http.StatusOK, "admin.categories.edit.gohtml", gin.H{
		"token":    csrf.GetToken(ctx),
		"category": handler.service.Find(id),
	})
}

func (handler CategoryHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	fmt.Println(id)
	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/categories/%s/edit", ctx.Param("id")), "Invalid Category ")
		return
	}
	CategoryRequest := &UpdateCategoryRequest{}
	err = ctx.ShouldBind(&CategoryRequest)
	if err != nil {

	}
	errs := validator.Validate(CategoryRequest)
	if errs != nil {
		fmt.Println(errs)
		helpers.RedirectFailedWithValidation(ctx, fmt.Sprintf("/admin/categories/%s/edit", ctx.Param("id")), errs, CategoryRequest)
		return
	}
	handler.service.Update(ctx, uint(id), CategoryRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/categories", "Category Updated successfully")

}
func (handler CategoryHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/categories", "Invalid Category ")
		return
	}
	handler.service.DeleteById(id)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/categories", "Deleted Successfully")

}

func (handler CategoryHandler) DeleteCategoryImage(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	handler.service.DeleteImageByCategoryId(id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Deleted Successfully",
	})

}

func NewCategoryHandler(DB *gorm.DB, commonData func(c *gin.Context, data gin.H) gin.H) *CategoryHandler {

	return &CategoryHandler{
		service:    NewCategoryService(DB),
		commonData: commonData,
	}
}

func withCommonData(c *gin.Context, data gin.H) gin.H {

	commonData := c.MustGet("CommonData").(gin.H)
	data["commonData"] = commonData

	return data
}
