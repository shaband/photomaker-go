package middlleware

import "github.com/gin-gonic/gin"

func LoadGlobleSiteMiddleware(r *gin.RouterGroup) {
	r.Use(commonDataMiddleware())
	r.Use(loadLocalizationMiddleware())
	r.Use(loadLangMiddleware())
}
