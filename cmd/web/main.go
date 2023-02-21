package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/templates"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()
	files := template.Must(template.New("").ParseFS(templates.FS, "site/layout/*.gohtml", "site/pages/*.gohtml", "*.gohtml"))
	router.SetHTMLTemplate(files)
	router.Static("assets", "./assets")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.gohtml", gin.H{})
	})
	//router.GET("/about", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "site/pages/about.gohtml", gin.H{})
	//})
	//
	//router.GET("/category", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "site/pages/category.gohtml", gin.H{})
	//})
	//
	//router.GET("/contact", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "site/pages/contact.gohtml", gin.H{})
	//})
	//
	//router.GET("/gallery", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "site/pages/gallery.gohtml", gin.H{})
	//})
	//
	//router.GET("/services", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "site/pages/services.gohtml", gin.H{})
	//})
	//
	//router.GET("/pill", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "site/pages/pill.gohtml", gin.H{})
	//})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
