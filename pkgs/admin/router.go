package admin

import (
	"fmt"
	"net/http"

	"github.com/shaband/photomaker-go/pkgs/modules/categories"
	"github.com/shaband/photomaker-go/pkgs/modules/contacts"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
	"github.com/shaband/photomaker-go/pkgs/modules/clients"
	"github.com/shaband/photomaker-go/pkgs/modules/services"
	"github.com/shaband/photomaker-go/pkgs/modules/sliders"
	"github.com/shaband/photomaker-go/pkgs/modules/users"
)

const userkey = "admin"

func Register(router *gin.RouterGroup) {

	middleware.LoadGlobalAdminMiddleware(router)
	authHandler := users.NewAuthHandler(database.GetConnection())

	guest := router.Group("/auth")
	guest.Use(IsGuest)
	// Login and logout routes
	guest.GET("/login", authHandler.GetLoginPage)
	guest.POST("/login", authHandler.Login)
	guest.GET("/logout", authHandler.Logout)
	// Private group, require authentication to access
	private := router.Group("/")
	private.Use(AuthRequired)
	// {
	// 	private.GET("/me", authHandler.Me)
	// 	private.GET("/status", authHandler.Status)
	// }

	categoriesHandler := categories.NewHandler(categories.NewService(database.GetConnection()), withCommonData)
	AddCrud(private.Group("/users"), users.NewHandler(database.GetConnection(), withCommonData))
	AddCrud(private.Group("/categories"), categoriesHandler)
	private.DELETE("/category-images/:id", categoriesHandler.DeleteCategoryImage)

	AddCrud(private.Group("/clients"), clients.NewHandler(clients.NewService(database.GetConnection()), withCommonData))
	AddCrud(private.Group("/sliders"), sliders.NewHandler(sliders.NewService(database.GetConnection()), withCommonData))
	AddCrud(private.Group("/services"), services.NewHandler(services.NewService(database.GetConnection()), withCommonData))

	settingsHandler := settings.NewHandler(settings.NewService(database.GetConnection()), withCommonData)
	contactsHandler := contacts.NewHandler(contacts.NewService(database.GetConnection()), withCommonData)
	private.GET("settings", settingsHandler.Index)
	private.GET("settings/:id/edit", settingsHandler.Edit)
	private.Match([]string{http.MethodPut, http.MethodPatch}, "settings/:id", settingsHandler.Update)

	private.GET("contacts", contactsHandler.Index)
	private.GET("contacts/:id/edit", contactsHandler.Edit)
	private.Match([]string{http.MethodPut, http.MethodPatch}, "contacts/:id", contactsHandler.Update)
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
		c.Redirect(http.StatusMovedPermanently, "/admin/auth/login")
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// AuthRequired is a simple middleware to check the session.
func IsGuest(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	fmt.Println(user)
	if user == nil {
		// Abort the request with the appropriate error code
		c.Next()
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/users")

	// Continue down the chain to handler etc
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
