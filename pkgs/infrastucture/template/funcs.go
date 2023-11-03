package template

import (
	"html/template"
	"time"
	//"github.com/Masterminds/sprig/v3"

	"github.com/shaband/photomaker-go/pkgs/site"
)

func GetOldWithDefault(inputs *map[string]interface{}, input string, defaultValue string) string {

	if inputs != nil {

		value, ok := (*inputs)[input]
		if ok {
			return value.(string)

		}

	}
	return defaultValue
}

var templateFuncs template.FuncMap = template.FuncMap{

	"getYear": func() int {
		return time.Now().Year()
	},
	"Trans": func(MessageID string) string {
		return site.Trans(MessageID, make(map[string]string))
	},

	"old": func(inputs *map[string]interface{}, input string) string {
		return GetOldWithDefault(inputs, input, "")
	},

	"oldWithDefault": GetOldWithDefault,
}
