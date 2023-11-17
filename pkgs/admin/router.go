package admin

import (
	"net/http"

	"github.com/shaband/photomaker-go/pkgs/modules/categories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
	"github.com/shaband/photomaker-go/pkgs/modules/users"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
)

const userkey = "user"

func Register(router *gin.RouterGroup) {

	middleware.LoadGlobalAdminMiddleware(router)
	authHandler := users.NewAuthHandler(database.GetConnection())

	categoriesHandler := categories.NewHandler(categories.NewService(database.GetConnection()), withCommonData)
	AddCrud(router.Group("/users"), users.NewHandler(database.GetConnection(), withCommonData))
	AddCrud(router.Group("/categories"), categoriesHandler)
	router.DELETE("/category-images/:id", categoriesHandler.DeleteCategoryImage)

	AddCrud(router.Group("/clients"), clients.NewHandler(clients.NewService(database.GetConnection()),withCommonData))
	AddCrud(router.Group("/sliders"), sliders.NewHandler(sliders.NewService(database.GetConnection()),withCommonData))
	AddCrud(router.Group("/services"), services.NewHandler(services.NewService(database.GetConnection()),withCommonData))

	guest := router.Group("/auth")
	// Login and logout routes
	guest.GET("/login", authHandler.GetLoginPage)
	guest.POST("/login", authHandler.Login)
	guest.GET("/logout", authHandler.Logout)
	// Private group, require authentication to access
	private := router.Group("/")
	private.Use(AuthRequired)
	{
		private.GET("/me", authHandler.Me)
		private.GET("/status", authHandler.Status)
	}

}

type CurdContract interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Store(ctx *gin.Context)
	Edit(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// AuthRequired is a simple middleware to check the session.
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func withCommonData(c *gin.Context, data gin.H) gin.H {

	commonData := c.MustGet("CommonData").(gin.H)
	data["commonData"] = commonData

	return data
}

func AddCrud(router *gin.RouterGroup, crud CurdContract) {
	router.GET("/", crud.Index)
	router.GET("/create", crud.Create)
	router.POST("/", crud.Store)
	router.GET("/:id/edit", crud.Edit)
	router.Match([]string{http.MethodPut, http.MethodPatch}, "/:id", crud.Update)
	router.DELETE("/:id", crud.Delete)
}
