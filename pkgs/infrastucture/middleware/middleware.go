package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoadGlobalSiteMiddleware(r *gin.RouterGroup) {

	r.Use(
		commonDataMiddleware(),
		loadLocalizationMiddleware(),
		loadLangMiddleware(),
	)
}
func LoadGlobalAdminMiddleware(r *gin.RouterGroup) {

	r.Use(
		commonDataMiddleware(),
		loadLocalizationMiddleware(),

		func(c *gin.Context) {
			{
				
				s := sessions.Default(c)
				var lang string = "en"
				s.Set(sessionKey, lang)
				AddToTemplateCommonData(c, commonDataLocaleKey, lang)
				c.Header("Accept-Language", lang)
				c.Request.Header.Set("Accept-Language", lang)
				c.Next()
			}
		},
		loadLangMiddleware(),
	)

}
