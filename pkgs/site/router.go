package site

import (
	"encoding/json"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"golang.org/x/text/language"
)

// func LoadSettings() gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		var results []*settings.Setting
// 		database.GetConnection().Select("slug", "value").Find(&results)

// 		c.Set("settings", results)

// 		// before request

// 		c.Next()

// 	}
// }

func SiteRegister(router *gin.RouterGroup) {

	// router.Use(LoadSettings())

	router.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		RootPath:         "pkgs/infrastructure/localize",
		AcceptLanguage:   []language.Tag{language.English, language.Arabic},
		DefaultLanguage:  language.Arabic,
		UnmarshalFunc:    json.Unmarshal,
		FormatBundleFile: "json",
	})))

	// apply i18n middleware
	//   router.Use(ginI18n.Localize(
	// 	ginI18n.WithGetLngHandle(
	// 	  func(context *gin.Context, defaultLng string) string {
	// 		lng := context.Query("lng")
	// 		if lng == "" {
	// 		  return defaultLng
	// 		}
	// 		return lng
	// 	  },
	// 	),
	//   ))

	handler := NewSiteHandler(database.GetConnection())

	router.GET("/", handler.IndexPage)
	router.GET("/about", handler.AboutPage)

	router.GET("/category/:id", handler.CategoryPage)

	router.GET("/contact", handler.ContactPage)

	router.GET("/gallery", handler.GalleryPage)

	router.GET("/services", handler.ServicesPage)

	router.GET("/pill", handler.PillPage)
}

// func trans(MessageID string, templateData map[string]string) string {
// 	return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
// 		MessageID:    MessageID,
// 		TemplateData: templateData,
// 	})
// }
