package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestSetGetRemoveCookie(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	SetCookie(c, "test", "value")
	value, err := GetCookie(c, "test")
	if err != nil || value != "value" {
		t.Errorf("SetCookie or GetCookie function failed")
	}

	RemoveCookie(c, "test")
	_, err = GetCookie(c, "test")
	if err == nil {
		t.Errorf("RemoveCookie function failed")
	}
}

func TestAbortError(t *testing.T) {
	err := errors.New("test error")
	AbortError(err)
}

func TestSecondMustBeNull(t *testing.T) {
	data := SecondMustBeNull("test", nil)
	if data != "test" {
		t.Errorf("SecondMustBeNull function failed")
	}
}

func TestFirstMustBeNull(t *testing.T) {
	data := FirstMustBeNull("test", nil)
	if data != "test" {
		t.Errorf("FirstMustBeNull function failed")
	}
}

func TestSaveFile(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	file, _ := c.FormFile("test")
	SaveFile(c, "dest", file)
}

func TestToJson(t *testing.T) {
	_, err := ToJson("test")
	if err != nil {
		t.Errorf("ToJson function failed")
	}
}

func TestToMap(t *testing.T) {
	m := toMap("{\"test\":\"value\"}")
	if (*m)["test"] != "value" {
		t.Errorf("toMap function failed")
	}
}

func TestRedirectSuccessWithMessage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	RedirectSuccessWithMessage(c, "url", "message")
}

func TestRedirectSuccessWithData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	RedirectSuccessWithData(c, "url", "message", "data")
}

func TestRedirectFailedWithMessage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	RedirectFailedWithMessage(c, "url", "message")
}

func TestRedirectFailedWithValidation(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	RedirectFailedWithValidation(c, "url", map[string]string{"error": "message"}, "inputs")
}
