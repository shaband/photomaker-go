package middleware

import (
	// "fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/database"
	"github.com/shaband/photomaker-go/pkgs/infrastucture/helpers"
	"github.com/shaband/photomaker-go/pkgs/modules/settings"
)

type commonData struct {
	gin.H
}

const commonDataKey = "CommonData"

func (c commonData) Set(key string, value any) {

	c.H[key] = value
}

var (
	once sync.Once

	commonDataInstance commonData
)

func ResolveCommonData() commonData {

	once.Do(func() {
		commonDataInstance = commonData{
			gin.H{
				"settings": settings.NewSettingService(database.GetConnection()).GetAllValuesPluckedBy("Slug"),
			},
		}
	})
	return commonDataInstance
}
func (c commonData) GetData() gin.H {
	return c.H
}

func commonDataMiddleware() gin.HandlerFunc {
	// Define common data to be added to the context for all requests

	// Add the common data to the context
	return func(c *gin.Context) {
		data := ResolveCommonData()
		data.H["currentPath"] = c.FullPath()
		loadOldData(c, &data)
		c.Set(commonDataKey, data.H)
		c.Next()
	}
}

func AddToCommonData[T any](key string, value T) gin.H {
	commonDataInstance.Set(key, value)
	return ResolveCommonData().H
}

func AddToTemplateCommonData[T any](c *gin.Context, key string, value T) {
	c.Set(commonDataKey, AddToCommonData(key, value))
}

func loadOldData(c *gin.Context, data *commonData) {

	errors, _ := helpers.GetCookieAsMap(c, helpers.ErrorsKey)
	helpers.RemoveCookie(c, helpers.ErrorsKey)

	inputs, _ := helpers.GetCookieAsMap(c, helpers.InputsKey)
	helpers.RemoveCookie(c, helpers.InputsKey)

	old_data, _ := helpers.GetCookieAsMap(c, helpers.DataKey)
	helpers.RemoveCookie(c, helpers.DataKey)

	status, _ := helpers.GetCookie(c, helpers.StatusKey)
	helpers.RemoveCookie(c, helpers.StatusKey)

	message, _ := helpers.GetCookie(c, helpers.MessageKey)
	helpers.RemoveCookie(c, helpers.MessageKey)

	data.H["__inputs"] = removeMapPrefix(inputs)
	data.H["__errors"] = removeMapPrefix(errors)
	data.H["__status"] = status
	data.H["__data"] = removeMapPrefix(old_data)
	data.H["__message"] = message
	// fmt.Println(data.H)
}

func ParsePointer(data *interface{}) interface{} {
	if data == nil {
		return nil
	}
	return *data
}

func removeMapPrefix(m *map[string]interface{}) *map[string]interface{} {
	newM := make(map[string]interface{})
	if m == nil {
		return m
	}
	for k, v := range *m {
		regx, _ := regexp.Compile(`[A-Za-z]*\.`)
		prefix := regx.FindString(k)
		if prefix != "" {
			k = strings.Trim(k, prefix)
		}
		newM[k] = v
	}
	return &newM
}
