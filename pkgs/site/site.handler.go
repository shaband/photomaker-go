package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
	csrf "github.com/utrack/gin-csrf"
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

	context.HTML(http.StatusOK, "site.pages.index.gohtml", handler.withCommonData(context, gin.H{
		"sliders": sliders.NewService(handler.db).GetSliders(),
	}))
}

func (handler *SiteHandler) AboutPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.pages.about.gohtml", handler.withCommonData(context, gin.H{
		"clients": clients.NewService(handler.db).All(),
	}))
}

func (handler *SiteHandler) CategoryPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.pages.category.gohtml", handler.withCommonData(context, gin.H{
		"category": categories.NewService(handler.db).GetSingleCategoryWithImages(context.Param("id")),
	}))
}

func (handler *SiteHandler) ContactPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.pages.contact.gohtml", handler.withCommonData(context, gin.H{

		"token":        csrf.GetToken(context),
		"serviceTypes": contacts.NewService(handler.db).GetAll(),
	}))
}

func (handler *SiteHandler) GalleryPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.pages.gallery.gohtml", handler.withCommonData(context, gin.H{
		"categories": categories.NewService(handler.db).All(),
	}))
}

func (handler *SiteHandler) SaveContactData(context *gin.Context) {

	form := contacts.NewService(handler.db).BindForm(context)
	err := validator.Validate(form)
	if err != nil {
		helpers.RedirectFailedWithValidation(context, "/contact", err, form)
		return
	} else {

		contacts.NewService(handler.db).SaveContactData(context, form)
		context.JSON(200, "success")
	}
}
func (handler *SiteHandler) ServicesPage(context *gin.Context) {

	context.HTML(http.StatusOK, "site.pages.services.gohtml", handler.withCommonData(context, gin.H{
		"services": services.NewService(handler.db).All(),
	}))
}
func (handler *SiteHandler) PillPage(context *gin.Context) {
	context.HTML(http.StatusOK, "site.pages.pill.gohtml", handler.withCommonData(context, gin.H{}))
}
