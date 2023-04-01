package site

import (
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
)

func LoadSettings() gin.HandlerFunc {

	return func(c *gin.Context) {

		var results []*settings.Setting
		database.GetConnection().Select("slug", "value").Find(&results)

		c.Set("settings", results)

		// before request

		c.Next()

	}
}

func SiteRegister(router *gin.RouterGroup) {

	router.Use(LoadSettings())

	handler := NewSiteHandler(database.GetConnection())

	router.GET("/", handler.IndexPage)
	router.GET("/about", handler.AboutPage)

	router.GET("/category/:id", handler.CategoryPage)

	router.GET("/contact", handler.ContactPage)

	router.GET("/gallery", handler.GalleryPage)

	router.GET("/services", handler.ServicesPage)

	router.GET("/pill", handler.PillPage)
}
