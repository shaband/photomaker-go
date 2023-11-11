package settings

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
	service    *Service
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler Handler) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin.settings.index.gohtml", withCommonData(ctx, gin.H{
		"token":    csrf.GetToken(ctx),
		"settings": handler.service.GetAll(),
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
		"user":  handler.service.Find(id),
	})
}

func (handler Handler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/settings/%s/edit", ctx.Param("id")), "Invalid Service ")
		return
	}
	SettingRequest := SettingRequest{}
	ctx.ShouldBind(&SettingRequest)

	errs := validator.Validate(SettingRequest)
	if errs != nil {
		fmt.Println(errs)

		helpers.RedirectFailedWithValidation(ctx, fmt.Sprintf("/admin/settings/%s/edit", ctx.Param("id")), errs, SettingRequest)
		return
	}
	handler.service.Update(uint(id), SettingRequest.ToEntity())
	helpers.RedirectSuccessWithMessage(ctx, "/admin/settings", "Service Updated successfully")

}
func NewHandler(DB *gorm.DB, commonData func(c *gin.Context, data gin.H) gin.H) *Handler {

	return &Handler{
		service:    NewService(DB),
		commonData: commonData,
	}
}

func withCommonData(c *gin.Context, data gin.H) gin.H {

	commonData := c.MustGet("CommonData").(gin.H)
	data["commonData"] = commonData

	return data
}

func (service *Service) Find(ID int) *Setting {
	setting := Setting{}
	service.db.Find(&setting, ID)
	return &setting
}
func (service Service) Update(ID uint, setting *Setting) {
	setting.ID = ID
	result := service.db.Save(setting)
	fmt.Println(result)
}

func (service *Service) DeleteById(ID int) *Service {
	service.db.Delete(&Service{}, ID)
	return nil
}
