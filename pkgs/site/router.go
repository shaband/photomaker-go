package site

import (
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
)

func SiteRegister(router *gin.RouterGroup) {

	middleware.LoadGlobalSiteMiddleware(router)

	handler := NewSiteHandler(database.GetConnection())

	router.GET("/", handler.IndexPage)
	router.GET("/about", handler.AboutPage)

	router.GET("/category/:id", handler.CategoryPage)

	router.GET("/contact", handler.ContactPage)
	router.POST("/contact", handler.SaveContactData)

	router.GET("/gallery", handler.GalleryPage)

	router.GET("/services", handler.ServicesPage)

	router.GET("/pill", handler.PillPage)

	router.GET("/admin/users", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin.users.index.gohtml", gin.H{})
	})
	router.GET("/admin/users/create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin.users.create.gohtml", gin.H{})
	})
	router.GET("/admin/categories", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin.categories.index.gohtml", gin.H{})
	})

}

func Trans(MessageID string, templateData map[string]string) string {
	return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID:    MessageID,
		TemplateData: templateData,
	})
}
