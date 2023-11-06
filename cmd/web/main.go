package main

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	// "github.com/gin-gonic/gin/binding"
	_ "github.com/joho/godotenv/autoload"
	"github.com/shaband/photomaker-go/pkgs/admin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/config"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/template"

	// "github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	method "github.com/bu/gin-method-override"
	"github.com/shaband/photomaker-go/pkgs/site"

	csrf "github.com/utrack/gin-csrf"
)

func main() {
	// binding.Validator = new(validator.DefaultValidator)
	router := gin.Default()
	router.Use(method.ProcessMethodOverride(router))

	gin.SetMode(os.Getenv("APP_MODE"))
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	database.Init(cfg)
	database.MakeMigration(database.GetConnection())

	middlewares(router)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.HTMLRender = template.LoadTemplates()
	router.Static("assets", "./assets")
	site.Register(router.Group("/"))
	admin.Register(router.Group("/admin"))
	err = router.Run(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080

}

func middlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
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
