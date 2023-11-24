package helpers

import (
	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, key string, value string) {

	c.SetCookie(key, string(value), 10, "/", c.Request.URL.Hostname(), false, true)

}

func GetCookie(c *gin.Context, key string) (string, error) {

	return c.Cookie(key)
}

func GetCookieAsMap(c *gin.Context, key string) (*map[string]interface{}, error) {

	str, err := GetCookie(c, key)
	if err != nil {
		return nil, err
	}
	return toMap(str), err
}

func RemoveCookie(c *gin.Context, key string) {
	c.SetCookie(key, "", -1, "/", c.Request.URL.Hostname(), false, true)

}
