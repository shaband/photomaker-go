package main

import (
	"path/filepath"
	"time"

	"text/template"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
	"github.com/shaband/photomaker-go/pkgs/site"
	csrf "github.com/utrack/gin-csrf"
)

func main() {
	router := gin.Default()

	database.Init()
	database.MakeMigration(database.GetConnection())
	loadmiddlewares(router)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.HTMLRender = loadTemplates("./templates")
	router.Static("assets", "./assets")
	site.SiteRegister(router.Group("/"))

	err := router.Run(":8081")
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
		r.AddFromFilesFuncs("site."+filepath.Base(sitePage), template.FuncMap{
			"getSetting": func(slug string) *string {
				Setting := settings.Setting{}
				database.GetConnection().Select("value").Where("slug = ?", slug).Find(&Setting)
				return &(Setting.Value)
			},
			"getYear": func() int {
				return time.Now().Year()
			},
			"Trans": func(MessageID string) string {
				return site.Trans(MessageID, make(map[string]string))
			},
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
func loadmiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))
}
