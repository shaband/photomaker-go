package clients

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
	service    ServiceInterface
	commonData func(c *gin.Context, data gin.H) gin.H
}

func (handler Handler) Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "admin.clients.index.gohtml", withCommonData(ctx, gin.H{
		"token":   csrf.GetToken(ctx),
		"clients": handler.service.GetAll(),
	}))

}

// func (handler Handler) Show(c *gin.Context)   {

// }
func (handler Handler) Create(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin.clients.create.gohtml", withCommonData(ctx, gin.H{
		"token": csrf.GetToken(ctx),
	}))
}
func (handler Handler) Store(ctx *gin.Context) {
	ClientRequest := CreateClientRequest{}
	_ = ctx.ShouldBind(&ClientRequest)

	errs := validator.Validate(ClientRequest)
	if errs != nil {
		helpers.RedirectFailedWithValidation(ctx, "/admin/clients/create", errs, ClientRequest)
		return
	}

	handler.service.Store(ctx, &ClientRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/clients", "Client created successfully")

}
func (handler Handler) Edit(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/clients", "Invalid Client ")
		return
	}
	ctx.HTML(http.StatusOK, "admin.clients.edit.gohtml", gin.H{
		"token":    csrf.GetToken(ctx),
		"category": handler.service.Find(id),
	})
}

func (handler Handler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	fmt.Println(id)
	if err != nil {
		helpers.RedirectFailedWithMessage(ctx, fmt.Sprintf("/admin/clients/%s/edit", ctx.Param("id")), "Invalid Client ")
		return
	}
	ClientRequest := &UpdateClientRequest{}
	err = ctx.ShouldBind(&ClientRequest)
	if err != nil {
		log.Println(err)
		return
	}
	errs := validator.Validate(ClientRequest)
	if errs != nil {
		fmt.Println(errs)
		helpers.RedirectFailedWithValidation(ctx, fmt.Sprintf("/admin/clients/%s/edit", ctx.Param("id")), errs, ClientRequest)
		return
	}
	handler.service.Update(ctx, uint(id), ClientRequest)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/clients", "Client Updated successfully")

}
func (handler Handler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		helpers.RedirectFailedWithMessage(ctx, "/admin/clients", "Invalid Client ")
		return
	}
	handler.service.DeleteById(id)
	helpers.RedirectSuccessWithMessage(ctx, "/admin/clients", "Deleted Successfully")

}

func (handler Handler) DeleteClientImage(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	handler.service.DeleteImageByClientId(id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Deleted Successfully",
	})

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
