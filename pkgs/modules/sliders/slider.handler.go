package sliders

import (
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
		"sliders": handler.service.GetAll(),
	}))

}

func (handler Handler) Create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.sliders.create.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
	}))
}
func (handler Handler) Store(ctx *gin.Context) {
	SliderRequest := CreateSliderRequest{}
	_ = ctx.ShouldBind(&SliderRequest)

	errs := validator.Validate(SliderRequest)
	if errs != nil {
		helpers.RedirectFailedWithValidation(ctx, "/admin/sliders/create", errs, SliderRequest)
		return
	}

	handler.service.Store(ctx, &SliderRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/sliders", "Slider created successfully")

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
