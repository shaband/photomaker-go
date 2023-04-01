package main

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/site"
)

func main() {
	router := gin.Default()

	database.Init()
	database.MakeMigration(database.GetConnection())

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.HTMLRender = loadTemplates("./templates")
	router.Static("assets", "./assets")
	site.SiteRegister(router.Group("/"))

	err := router.Run(":8080")
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
		r.AddFromFilesFuncs("site."+filepath.Base(sitePage), template.FuncMap{
			"getSetting":func (name)  {
					database.GetConnection().Select("value").Find()
			}
		}, files...)

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
