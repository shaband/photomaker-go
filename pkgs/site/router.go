package site

import (
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
	"github.com/gin-contrib/sessions"
)

func SiteRegister(router *gin.RouterGroup) {

	middleware.LoadGlobalSiteMiddleware(router)

	handler := NewSiteHandler(database.GetConnection())

	router.GET("/", handler.IndexPage)
	router.GET("/about", handler.AboutPage)

	router.GET("/category/:id", handler.CategoryPage)

	router.GET("/contact", handler.ContactPage)
	router.POST("/contact", handler.SaveContactData)

	router.GET("/gallery", handler.GalleryPage)

	router.GET("/services", handler.ServicesPage)

	router.GET("/pill", handler.PillPage)

	router.GET("/admin/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin.auth.login.gohtml", gin.H{})
	})
	

	authMiddleware := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			session := sessions.Default(c)
			user := session.Get("user")

			if user == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
				c.Abort()
				return
			}

			c.Next()
		}
	}

	// Protected route that requires authentication
	router.GET("/protected", authMiddleware(), func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		c.JSON(http.StatusOK, gin.H{"message": "Authenticated", "user": user})
	})

	// Login route
	router.POST("admin/login", func(c *gin.Context) {
		session := sessions.Default(c)
		user := c.PostForm("user")

		// Perform authentication logic, e.g., check credentials against a database

		// Simulating a successful authentication
		session.Set("user", user)
		session.Save()

		c.JSON(http.StatusOK, gin.H{"message": "Logged in", "user": user})
	})

	// Logout route
	router.POST("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		
		session.Clear()
		session.Save()

		c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
	})
}

func Trans(MessageID string, templateData map[string]string) string {
	return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID:    MessageID,
		TemplateData: templateData,
	})
}
