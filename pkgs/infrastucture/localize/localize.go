package localize

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLoadLocalizationMiddleware(t *testing.T) {
	// Create a new gin.Context instance
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Call the LoadLocalizationMiddleware function with the context
	LoadLocalizationMiddleware(context)

	// Assert that the middleware is correctly applied and the expected changes are made to the context
	// Add your assertions here
}
