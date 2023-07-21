package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
	"net/http"
)

const userkey = "user"

func AdminRegister(router *gin.RouterGroup) {

	middleware.LoadGlobalSiteMiddleware(router)

	authHandler := newAuthHandler(database.GetConnection())

	router.GET("users", func(c *gin.Context) {

		c.HTML(http.StatusOK, "admin.users.index.gohtml", gin.H{})
	})
	guest := router.Group("/auth")
	// Login and logout routes
	guest.GET("/login", authHandler.GetLoginPage)
	guest.POST("/login", authHandler.Login)
	guest.GET("/logout", authHandler.Logout)
	// Private group, require authentication to access
	private := router.Group("/admin")
	private.Use(AuthRequired)
	{
		private.GET("/me", authHandler.Me)
		private.GET("/status", authHandler.Status)
	}

}

type CurdContract interface {
	Index()
	Create()
	Store()
	Edit()
	Update()
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
