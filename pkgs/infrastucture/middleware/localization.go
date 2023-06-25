package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	ginI18n "github.com/gin-contrib/i18n"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"golang.org/x/text/language"
)

const sessionKey = "lang"
const commonDataLocaleKey = "locale"

var AcceptLanguage []language.Tag = []language.Tag{language.English, language.Arabic}
var DefaultLanguage = language.Arabic
var CurrentLanguage= "ar"

var localizationPath string = "pkgs/infrastucture/localize"

func loadLangMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		{
			s := sessions.Default(c)
			var lang string = "ar"
			if s.Get(sessionKey) != nil {
				lang = s.Get(sessionKey).(string)
			}
			lang = c.DefaultQuery(sessionKey, lang)

			s.Set(sessionKey, lang)
			err := s.Save()
			if err != nil {
				fmt.Println(err)
			}
			AddToTemplateCommonData(c, commonDataLocaleKey, lang)
			c.Header("Accept-Language", lang)
			c.Request.Header.Set("Accept-Language", lang)
			CurrentLanguage=lang
			c.Next()
		}
	}
}

func loadLocalizationMiddleware() gin.HandlerFunc {
	return ginI18n.Localize(
		ginI18n.WithBundle(
			&ginI18n.BundleCfg{
				RootPath:         localizationPath,
				AcceptLanguage:   AcceptLanguage,
				DefaultLanguage:  DefaultLanguage,
				UnmarshalFunc:    json.Unmarshal,
				FormatBundleFile: "json",
			}),

		ginI18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				s := sessions.Default(context)
				if s.Get(sessionKey) != nil {
					defaultLng = s.Get(sessionKey).(string)
				}
				return defaultLng
			},
		),
	)
}

func Trans(MessageID string, templateData map[string]string) string {
	return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID:    MessageID,
		TemplateData: templateData,
	})
}
