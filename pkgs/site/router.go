package site

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
)

func Register(router *gin.RouterGroup) {

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

	// router.GET("/admin/login", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "admin.auth.login.gohtml", gin.H{})
	// })

}

func Trans(MessageID string, templateData map[string]string) string {
	return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID:    MessageID,
		TemplateData: templateData,
	})
}
