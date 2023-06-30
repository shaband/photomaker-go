package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func SaveFile(c *gin.Context, dest string, file *multipart.FileHeader) string {
	AbortError(c.SaveUploadedFile(file, fmt.Sprintf(
		dest,
		"/",
		file.Filename)))
	return dest
}

func ToJson(val any) ([]byte, error) {

	bytes, err := json.Marshal(val)

	return bytes, err

}

func toMap(s string) *map[string]interface{} {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		fmt.Println(err)
	}
	return &m
}
