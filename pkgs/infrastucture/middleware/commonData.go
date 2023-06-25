package middleware

import (
	"fmt"
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
		errors, _ := helpers.GetCookieAsMap(c, "errors")
		inputs, _ := helpers.GetCookieAsMap(c, "inputs")
		helpers.RemoveCookie(c, "errors")
		helpers.RemoveCookie(c, "inputs")
		data.H["currentPath"] = c.FullPath()
		data.H["old_inputs"] = inputs
		data.H["errors"] = errors

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
