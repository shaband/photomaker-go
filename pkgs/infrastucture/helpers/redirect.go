package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const ErrorCode = "error"
const SuccessCode = "success"


const InputsKey = "inputs"
const StatusKey = "status"
const MessageKey = "message"
const DataKey = "data"
const ErrorsKey = "errors"

func GetAllKeys() [5]string {

	return [5]string{
		StatusKey,
		MessageKey,
		ErrorCode,
		InputsKey,
		DataKey,
	}
}

func RedirectSuccessWithMessage(c *gin.Context, url string, message string) {

	SetCookie(c, StatusKey, SuccessCode)
	SetCookie(c, MessageKey, message)
	c.Redirect(http.StatusMovedPermanently, url)
}

func RedirectSuccessWithData[T any](c *gin.Context, url string, message string, data T) {

	errJson := SecondMustBeNull(ToJson(data))
	SetCookie(c, ErrorsKey, string(errJson))
	RedirectSuccessWithMessage(c, url, message)
}

func RedirectFailedWithMessage(c *gin.Context, url string, message string) {

	SetCookie(c, StatusKey, ErrorCode)
	SetCookie(c, MessageKey, message)
	c.Redirect(http.StatusMovedPermanently, url)
}

func RedirectFailedWithValidation[T any](c *gin.Context, url string, errors map[string]string, inputs T) {

	errJson := SecondMustBeNull(ToJson(errors))
	SetCookie(c, ErrorsKey, string(errJson))
	inputsJson := SecondMustBeNull(ToJson(inputs))
	SetCookie(c, InputsKey, string(inputsJson))
	RedirectFailedWithMessage(c, url, "Invalid Inputs")
}
