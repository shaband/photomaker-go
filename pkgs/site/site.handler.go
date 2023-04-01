package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
	"gorm.io/gorm"
)

type SiteHandler struct {
	db *gorm.DB
}

func NewSiteHandler(db *gorm.DB) *SiteHandler {
	return &SiteHandler{db: db}

}

func (handler *SiteHandler) IndexPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.index.gohtml", gin.H{
		"settings": context.MustGet("settings").([]*settings.Setting),
		"sliders":  sliders.NewSliderService(handler.db).GetSliders(),
	})
}

func (handler *SiteHandler) AboutPage(context *gin.Context) {

	Clients := []*clients.Client{}

	handler.db.Find(&Clients)

	context.HTML(http.StatusOK, "site.about.gohtml", gin.H{
		"settings": context.MustGet("settings").([]*settings.Setting),
		"clients":  Clients,
	})
}

func (handler *SiteHandler) CategoryPage(context *gin.Context) {
	Category := categories.Category{}
	handler.db.Preload("Images").First(&Category, context.Param("id"))

	context.HTML(http.StatusOK, "site.category.gohtml", gin.H{
		"settings": context.MustGet("settings").([]*settings.Setting),
		"category": Category,
	})
}
func (handler *SiteHandler) ContactPage(context *gin.Context) {
	ServiceTypes := contacts.ServiceType{}
	handler.db.Preload("Items").Find(&ServiceTypes)

	context.HTML(http.StatusOK, "site.contact.gohtml", gin.H{
		"settings":     context.MustGet("settings").([]*settings.Setting),
		"serviceTypes": contacts.NewContractService(handler.db).GetAll(),
	})
}
func (handler *SiteHandler) GalleryPage(context *gin.Context) {
	Categories := []*categories.Category{}
	handler.db.Find(&Categories)

	context.HTML(http.StatusOK, "site.gallery.gohtml", gin.H{
		"settings":   context.MustGet("settings").([]*settings.Setting),
		"categories": Categories,
	})
}
func (handler *SiteHandler) ServicesPage(context *gin.Context) {
	Services := []*services.Service{}
	handler.db.Find(&Services)

	context.HTML(http.StatusOK, "site.services.gohtml", gin.H{
		"settings": context.MustGet("settings").([]*settings.Setting),
		"services": Services,
	})
}
func (handler *SiteHandler) PillPage(context *gin.Context) {
	context.HTML(http.StatusOK, "site.pill.gohtml", gin.H{
		"settings": context.MustGet("settings").([]*settings.Setting),
	})
}
