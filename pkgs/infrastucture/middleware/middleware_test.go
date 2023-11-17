package middleware

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetGetData(t *testing.T) {
	data := commonData{gin.H{}}
	data.Set("key", "value")
	assert.Equal(t, "value", data.GetData()["key"])
}

func TestAddToCommonData(t *testing.T) {
	data := AddToCommonData("key", "value")
	assert.Equal(t, "value", data["key"])
}

func TestAddToTemplateCommonData(t *testing.T) {
	c, _ := gin.CreateTestContext(nil)
	AddToTemplateCommonData(c, "key", "value")
	value, _ := c.Get(commonDataKey)
	assert.Equal(t, "value", value.(gin.H)["key"])
}

func TestLoadOldData(t *testing.T) {
	// TODO: Implement this test
}

func TestLoadLangMiddleware(t *testing.T) {
	// TODO: Implement this test
}

func TestLoadLocalizationMiddleware(t *testing.T) {
	// TODO: Implement this test
}

func TestTrans(t *testing.T) {
	// TODO: Implement this test
}

func TestErrorHandler(t *testing.T) {
	// TODO: Implement this test
}

func TestLoadGlobalSiteMiddleware(t *testing.T) {
	// TODO: Implement this test
}

func TestLoadGlobalAdminMiddleware(t *testing.T) {
	// TODO: Implement this test
}
