package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/templates"
	"html/template"
	"net/http"
)

func main() {
	router := gin.Default()
	files := template.Must(template.New("").ParseFS(templates.FS, "*.gohtml", "site/layout/*.gohtml", "site/pages/*.gohtml"))
	router.SetHTMLTemplate(files)
	router.Static("assets", "./assets")
	router.GET("/", func(context *gin.Context) {

		context.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "Posts",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
