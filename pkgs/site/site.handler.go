package site

import (
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"

	// csrf "github.com/utrack/gin-csrf"
	"gorm.io/gorm"
)

type SiteHandler struct {
	db *gorm.DB
}

func (SiteHandler) withCommonData(c *gin.Context, data gin.H) gin.H {

	commonData := c.MustGet("CommonData").(gin.H)
	data["commonData"] = commonData

	return data
}

func NewSiteHandler(db *gorm.DB) *SiteHandler {
	return &SiteHandler{db: db}

}

func (handler *SiteHandler) IndexPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.index.gohtml", handler.withCommonData(context, gin.H{
		"sliders": sliders.NewSliderService(handler.db).GetSliders(),
	}))
}

func (handler *SiteHandler) AboutPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.about.gohtml", handler.withCommonData(context, gin.H{
		"clients": clients.NewClientSerive(handler.db).All(),
	}))
}

func (handler *SiteHandler) CategoryPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.category.gohtml", handler.withCommonData(context, gin.H{
		"category": categories.NewContractService(handler.db).GetSingleCategoryWithImages(context.Param("id")),
	}))
}

func (handler *SiteHandler) ContactPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.contact.gohtml", handler.withCommonData(context, gin.H{

		"token":        "aaaaa",
		"serviceTypes": contacts.NewContractService(handler.db).GetAll(),
	}))
}

func (handler *SiteHandler) GalleryPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.gallery.gohtml", handler.withCommonData(context, gin.H{
		"categories": categories.NewContractService(handler.db).All(),
	}))
}

func (hanlder *SiteHandler) SaveContactData(context *gin.Context) {

	var form contacts.ContactForm

	if err := context.ShouldBind(&form); err != nil {
		// handle error
		return
	}

	context.JSON(200, form)
}
func (handler *SiteHandler) ServicesPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.services.gohtml", handler.withCommonData(context, gin.H{
		"services": services.NewServicesSerive(handler.db).All(),
	}))
}
func (handler *SiteHandler) PillPage(context *gin.Context) {
	context.HTML(http.StatusOK, "site.pill.gohtml", handler.withCommonData(context, gin.H{}))
}
