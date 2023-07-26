package admin

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/middleware"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/validator"
	"github.com/shaband/photomaker-go/pkgs/modules/users"
	csrf "github.com/utrack/gin-csrf"
)

const userkey = "user"

func AdminRegister(router *gin.RouterGroup) {

	middleware.LoadGlobalSiteMiddleware(router)

	authHandler := newAuthHandler(database.GetConnection())

	router.GET("users", func(c *gin.Context) {

		
		c.HTML(http.StatusOK, "admin.users.index.gohtml", withCommonData(c, gin.H{
			"users": users.NewUserService(database.GetConnection()).GetAll(),
		}))
	})
	router.GET("/users/create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin.users.create.gohtml", withCommonData(ctx, gin.H{
			"token": csrf.GetToken(ctx),
		}))
	})
	router.POST("/users", func(ctx *gin.Context) {
		UserRequest := users.UserRequest{}
		ctx.ShouldBind(&UserRequest)
		errs := validator.Validate(UserRequest)
		if errs != nil {

			helpers.RedirectFailedWithValidation(ctx, "/admin/users/create", errs, UserRequest)
			return
		}
		user := UserRequest.ToEntity()
		database.GetConnection().Create(user)
		helpers.RedirectSuccessWithMessage(ctx, "/admin/users", "User created successfully")

	})

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

func withCommonData(c *gin.Context, data gin.H) gin.H {

	commonData := c.MustGet("CommonData").(gin.H)
	data["commonData"] = commonData

	return data
}
