package template

import (
	"github.com/shaband/photomaker-go/pkgs/site"
	"html/template"
	"time"
)

var templateFuncs template.FuncMap = template.FuncMap{

	"getYear": func() int {
		return time.Now().Year()
	},
	"Trans": func(MessageID string) string {
		return site.Trans(MessageID, make(map[string]string))
	}}
