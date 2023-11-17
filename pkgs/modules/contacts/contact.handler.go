package contacts

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	csrf "github.com/utrack/gin-csrf"
)

type Handler struct {
	service    Service
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler Handler) Edit(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/contacts", "Invalid Contact ")
		return
	}
	ctx.HTML(http.StatusOK, "admin.contacts.edit.gohtml", gin.H{
		"token":   csrf.GetToken(ctx),
		"contact": handler.service.Find(id),
	})
}

func (handler Handler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	fmt.Println(id)
	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/contacts/%s/edit", ctx.Param("id")), "Invalid Contact ")
		return
	}
	ContactRequest := &ContactRequest{}
	err = ctx.ShouldBind(&ContactRequest)
	if err != nil {
		log.Println(err)
		return
	}
	errs := validator.Validate(ContactRequest)
	if errs != nil {
		fmt.Println(errs)
		helpers.RedirectFailedWithValidation(ctx, fmt.Sprintf("/admin/contacts/%s/edit", ctx.Param("id")), errs, ContactRequest)
		return
	}
	handler.service.Update(ctx, uint(id), ContactRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/contacts", "Contact Updated successfully")

}

func NewHandler(service Service, commonData func(c *gin.Context, data gin.H) gin.H) *Handler {

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
