package sliders

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	csrf "github.com/utrack/gin-csrf"
)

type Handler struct {
	service    ServiceInterface
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler Handler) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin.sliders.index.gohtml", withCommonData(ctx, gin.H{
		"token":   csrf.GetToken(ctx),
		"sliders": handler.service.All(),
	}))

}

func (handler Handler) Create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.sliders.create.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
	}))
}
func (handler Handler) Store(ctx *gin.Context) {
	SliderRequest := SliderRequest{}
	_ = ctx.ShouldBind(&SliderRequest)

	errs := validator.Validate(SliderRequest)
	if errs != nil {
		helpers.RedirectFailedWithValidation(ctx, "/admin/sliders/create", errs, SliderRequest)
		return
	}

	handler.service.Store(ctx, &SliderRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/sliders", "Slider created successfully")

}
func (handler Handler) Edit(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/services", "Invalid Service ")
		return
	}
	ctx.HTML(http.StatusOK, "admin.services.edit.gohtml", gin.H{
		"token": csrf.GetToken(ctx),
		"user":  handler.service.Find(id),
	})
}

func (handler Handler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/services/%s/edit", ctx.Param("id")), "Invalid Service ")
		return
	}
	Request := SliderRequest{}
	ctx.ShouldBind(&Request)

	errs := validator.Validate(Request)
	if errs != nil {
		fmt.Println(errs)

		helpers.RedirectFailedWithValidation(ctx, fmt.Sprintf("/admin/services/%s/edit", ctx.Param("id")), errs, Request)
		return
	}
	handler.service.Update(ctx,uint(id),&Request)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/services", "Service Updated successfully")

}
func (handler Handler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/sliders", "Invalid Slider ")
		return
	}
	handler.service.DeleteById(id)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/sliders", "Deleted Successfully")

}

func NewHandler(service ServiceInterface, commonData func(c *gin.Context, data gin.H) gin.H) *Handler {

	return &Handler{
		service:    service,
		commonData: commonData,
	}
}

func withCommonData(c *gin.Context, data gin.H) gin.H {

	commonData := c.MustGet("CommonData").(gin.H)
	data["commonData"] = commonData

	return data
}
