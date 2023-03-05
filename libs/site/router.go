package site

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func SiteRegister(router *gin.RouterGroup){

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "site.index.gohtml", gin.H{})
	})
	router.GET("/about", func(context *gin.Context) {
		context.HTML(http.StatusOK, "site.about.gohtml", gin.H{})
	})

	router.GET("/category", func(context *gin.Context) {
		context.HTML(http.StatusOK, "site.category.gohtml", gin.H{})
	})

	router.GET("/contact", func(context *gin.Context) {
		context.HTML(http.StatusOK, "site.contact.gohtml", gin.H{})
	})

	router.GET("/gallery", func(context *gin.Context) {
		context.HTML(http.StatusOK, "site.gallery.gohtml", gin.H{})
	})

	router.GET("/services", func(context *gin.Context) {
		context.HTML(http.StatusOK, "site.services.gohtml", gin.H{})
	})

	router.GET("/pill", func(context *gin.Context) {
		context.HTML(http.StatusOK, "site.pill.gohtml", gin.H{})
	})
}


