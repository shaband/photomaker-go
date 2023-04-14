package site

import (
	"encoding/json"
	"fmt"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-contrib/sessions"

	// "github.com/gin-contrib/sessions"

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
	// router.Use(LoadLang())

	router.Use(loadLocalization())

	handler := NewSiteHandler(database.GetConnection())

	router.GET("/", handler.IndexPage)
	router.GET("/about", handler.AboutPage)

	router.GET("/category/:id", handler.CategoryPage)

	router.GET("/contact", handler.ContactPage)

	router.GET("/gallery", handler.GalleryPage)

	router.GET("/services", handler.ServicesPage)

	router.GET("/pill", handler.PillPage)
}

func Trans(MessageID string, templateData map[string]string) string {
	return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID:    MessageID,
		TemplateData: templateData,
	})
}

func LoadLang() gin.HandlerFunc {
	return func(c *gin.Context) {
		{
			s := sessions.Default(c)
			var lang string = "ar"
			if s.Get("lang") != nil || s.Get("lang") != "" {
				lang = s.Get("lang").(string)
				fmt.Println("From Session =======================================")
				fmt.Println(lang)
			}
			lang = c.DefaultQuery("lang", lang)
			fmt.Print("final==============================================+")
			fmt.Println(lang)
			s.Set("lang", lang)
			err := s.Save()
			fmt.Println(err)
			fmt.Print("from session  2 ===>   ")
			fmt.Println(lang)
			c.Next()
		}
	}
}

func loadLocalization() gin.HandlerFunc {
	return ginI18n.Localize(
		ginI18n.WithBundle(
			&ginI18n.BundleCfg{
				RootPath:         "pkgs/infrastucture/localize",
				AcceptLanguage:   []language.Tag{language.English, language.Arabic},
				DefaultLanguage:  language.Arabic,
				UnmarshalFunc:    json.Unmarshal,
				FormatBundleFile: "json",
			}),

		ginI18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				s := sessions.Default(context)
				if s.Get("lang") != nil {
					defaultLng = s.Get("lang").(string)
				}
				return defaultLng
			},
		),
	)
}
