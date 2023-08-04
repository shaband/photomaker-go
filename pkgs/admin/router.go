package admin

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
	"github.com/shaband/photomaker-go/pkgs/modules/users"
)

const userkey = "user"

func AdminRegister(router *gin.RouterGroup) {

	userHandler := users.NewUserHandler(database.GetConnection(), withCommonData)
	middleware.LoadGlobalAdminMiddleware(router)

	authHandler := newAuthHandler(database.GetConnection())

	router.GET("users", userHandler.Index)
	router.GET("/users/create", userHandler.Create)
	router.POST("/users", userHandler.Store)
	router.GET("/users/:user/edit", userHandler.Edit)
	router.Match([]string{http.MethodPut, http.MethodPatch}, "/users/:user", userHandler.Update)
	router.DELETE("/users/:user", userHandler.Delete)

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
