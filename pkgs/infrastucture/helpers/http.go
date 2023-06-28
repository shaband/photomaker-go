package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectWithErrorsAndInputs[T any](c *gin.Context, url string, errors map[string]string, inputs T) {
	errJson, _ := ToJson(errors)
	fmt.Println(string(errJson))
	SetCookie(c, "errors", string(errJson))

	inputsJson, _ := ToJson(inputs)
	SetCookie(c, "inputs", string(inputsJson))

	fmt.Println(string(errJson), string(inputsJson))
	c.Redirect(http.StatusMovedPermanently, url)

}
