package settings

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	csrf "github.com/utrack/gin-csrf"
	
)

type Handler struct {
	service    ServiceInterface
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler Handler) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin.settings.index.gohtml", withCommonData(ctx, gin.H{
		"token":    csrf.GetToken(ctx),
		"settings": handler.service.All(),
	}))

}

func (handler Handler) Edit(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/settings", "Invalid Service ")
		return
	}
	ctx.HTML(http.StatusOK, "admin.settings.edit.gohtml", gin.H{
		"token": csrf.GetToken(ctx),
		"setting":  handler.service.Find(id),
	})
}

func (handler Handler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/settings/%s/edit", ctx.Param("id")), "Invalid Service ")
		return
	}
	// ctx.ShouldBind(&SettingRequest)

	// errs := validator.Validate(SettingRequest)
	// if errs != nil {
		// fmt.Println(errs)

		// helpers.RedirectFailedWithValidation(ctx, fmt.Sprintf("/admin/settings/%s/edit", ctx.Param("id")), errs, SettingRequest)
		// return
	// }
err,_	=handler.service.Update(ctx,uint(id))

	if err != nil {
		fmt.Println(err)
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/settings/%s/edit", ctx.Param("id")), err.Error())
		return
	
	}
	helpers.RedirectSuccessWithMessage(ctx, "/admin/settings", "Service Updated successfully")

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
