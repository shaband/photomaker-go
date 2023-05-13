package middleware


import "github.com/gin-gonic/gin"

func LoadGlobalSiteMiddleware(r *gin.RouterGroup) {
	
	r.Use(commonDataMiddleware())
	r.Use(loadLocalizationMiddleware())
	r.Use(loadLangMiddleware())
}
