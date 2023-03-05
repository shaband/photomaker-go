package main

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/libs/site"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	_, err := gorm.Open(sqlite.Open("./database/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.HTMLRender = loadTemplates("./templates")
	router.Static("assets", "./assets")
	site.SiteRegister(router.Group("/"))
	// router.GET("/", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "site.index.gohtml", gin.H{})
	// })
	// router.GET("/about", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "site.about.gohtml", gin.H{})
	// })

	// router.GET("/category", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "site.category.gohtml", gin.H{})
	// })

	// router.GET("/contact", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "site.contact.gohtml", gin.H{})
	// })

	// router.GET("/gallery", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "site.gallery.gohtml", gin.H{})
	// })

	// router.GET("/services", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "site.services.gohtml", gin.H{})
	// })

	// router.GET("/pill", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "site.pill.gohtml", gin.H{})
	// })

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	siteLayouts, err := filepath.Glob(templatesDir + "/site/layout/*.gohtml")
	if err != nil {
		panic(err.Error())
	}

	sitePages, err := filepath.Glob(templatesDir + "/site/pages/*.gohtml")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our articleLayouts/ and articles/ directories
	for _, sitePage := range sitePages {
		siteLayoutCopy := make([]string, len(siteLayouts))
		copy(siteLayoutCopy, siteLayouts)
		files := append(siteLayoutCopy, sitePage)
		fmt.Println(files)
		r.AddFromFiles("site."+filepath.Base(sitePage), files...)
	}
	// adminLayouts, err := filepath.Glob(templatesDir + "/layouts/admin/*.gohtml")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// admins, err := filepath.Glob(templatesDir + "/admins/*.html")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // Generate our templates map from our adminLayouts/ and admins/ directories
	// for _, admin := range admins {
	// 	layoutCopy := make([]string, len(adminLayouts))
	// 	copy(layoutCopy, adminLayouts)
	// 	files := append(layoutCopy, admin)
	// 	r.AddFromFiles(filepath.Base(admin), files...)
	// }
	return r
}

// func createTemplateNameFormPath(templatePath string) string {

// 	fmt.Println(filepath.Clean(templatePath))
// 	// template_without_prefix := strings.Replace(templatePath, "templates/", "", 1)
// 	template_with_suffix := strings.Trim(templatePath, ".gohtml")
// 	// fmt.Println(template_with_suffix)
// 	return strings.ReplaceAll(template_with_suffix, "/", ".")

// }
