package template

import (
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
	"github.com/shaband/photomaker-go/pkgs/site"
	"html/template"
	"time"
)

var templateFuncs template.FuncMap = template.FuncMap{
	"getSetting": func(slug string) *string {
		return settings.NewSettingService(database.GetConnection()).FindValue(slug)
	},
	"getYear": func() int {
		return time.Now().Year()
	},
	"Trans": func(MessageID string) string {
		return site.Trans(MessageID, make(map[string]string))
	}}
