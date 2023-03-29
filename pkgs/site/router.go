package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"

	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
)

func LoadSettings() gin.HandlerFunc {

	return func(c *gin.Context) {

		var results []*settings.Setting
		database.GetConnection().Find(&results)

		c.Set("settings", results)

		// before request

		c.Next()

	}
}

func SiteRegister(router *gin.RouterGroup) {

	router.Use(LoadSettings())
	router.GET("/", func(context *gin.Context) {

		context.HTML(http.StatusOK, "site.index.gohtml", gin.H{
			"settings": context.MustGet("settings").([]*settings.Setting),
			"sliders":  sliders.NewSliderService(database.GetConnection()).GetSliders(),
		})
	})
	router.GET("/about", func(context *gin.Context) {

		Clients := []*clients.Client{}

		database.GetConnection().Find(&Clients)

		context.HTML(http.StatusOK, "site.about.gohtml", gin.H{
			"settings": context.MustGet("settings").([]*settings.Setting),
			"clients":  Clients,
		})
	})

	router.GET("/category/:id", func(context *gin.Context) {
		Category := categories.Category{}
		database.GetConnection().Preload("Images").First(&Category, context.Param("id"))

		context.HTML(http.StatusOK, "site.category.gohtml", gin.H{
			"settings": context.MustGet("settings").([]*settings.Setting),
			"category": Category,
		})
	})

	router.GET("/contact", func(context *gin.Context) {

		ServiceTypes := contacts.ServiceType{}
		database.GetConnection().Preload("Items").Find(&ServiceTypes)

		context.HTML(http.StatusOK, "site.contact.gohtml", gin.H{
			"settings":     context.MustGet("settings").([]*settings.Setting),
			"serviceTypes": ServiceTypes,
		})
	})

	router.GET("/gallery", func(context *gin.Context) {
		Categories := []*categories.Category{}
		database.GetConnection().Find(&Categories)

		context.HTML(http.StatusOK, "site.gallery.gohtml", gin.H{
			"settings":   context.MustGet("settings").([]*settings.Setting),
			"categories": Categories,
		})
	})

	router.GET("/services", func(context *gin.Context) {

		Services := []*services.Service{}
		database.GetConnection().Find(&Services)

		context.HTML(http.StatusOK, "site.services.gohtml", gin.H{
			"settings": context.MustGet("settings").([]*settings.Setting),
			"services": Services,
		})
	})

	router.GET("/pill", func(context *gin.Context) {
		context.HTML(http.StatusOK, "site.pill.gohtml", gin.H{
			"settings": context.MustGet("settings").([]*settings.Setting),
		})
	})
}
