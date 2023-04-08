package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
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
		"sliders": sliders.NewSliderService(handler.db).GetSliders(),
	})
}

func (handler *SiteHandler) AboutPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.about.gohtml", gin.H{
		"clients": clients.NewClientSerive(handler.db).All(),
	})
}

func (handler *SiteHandler) CategoryPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.category.gohtml", gin.H{
		"category": categories.NewContractService(handler.db).GetSingleCategoryWithImages(context.Param("id")),
	})
}

func (handler *SiteHandler) ContactPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.contact.gohtml", gin.H{
		"serviceTypes": contacts.NewContractService(handler.db).GetAll(),
	})
}

func (handler *SiteHandler) GalleryPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.gallery.gohtml", gin.H{
		"categories": categories.NewContractService(handler.db).All(),
	})
}

func (handler *SiteHandler) ServicesPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.services.gohtml", gin.H{
		"services": services.NewServicesSerive(handler.db).All(),
	})
}
func (handler *SiteHandler) PillPage(context *gin.Context) {
	context.HTML(http.StatusOK, "site.pill.gohtml", gin.H{})
}
