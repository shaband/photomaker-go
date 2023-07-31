package main

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	// "github.com/gin-gonic/gin/binding"
	_ "github.com/joho/godotenv/autoload"
	"github.com/shaband/photomaker-go/pkgs/admin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/template"

	// "github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	"github.com/shaband/photomaker-go/pkgs/site"
	csrf "github.com/utrack/gin-csrf"
)

func main() {
	// binding.Validator = new(validator.DefaultValidator)
	router := gin.Default()
	gin.SetMode(os.Getenv("APP_MODE"))
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	database.Init()
	database.MakeMigration(database.GetConnection())

	loadmiddlewares(router)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.HTMLRender = template.LoadTemplates()
	router.Static("assets", "./assets")
	site.SiteRegister(router.Group("/"))
	admin.AdminRegister(router.Group("/admin"))
	err := router.Run(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080

}

func loadmiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	// router.Use(gin.Recovery())
	router.Use(middleware.ErrorHandler())
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	router.Use(sessions.Sessions(os.Getenv("SESSION_NAME"), store))
	router.Use(csrf.Middleware(csrf.Options{
		Secret: os.Getenv("CSRF_SECRET"),
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

}
