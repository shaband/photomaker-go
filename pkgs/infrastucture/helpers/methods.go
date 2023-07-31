package helpers

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveFile(c *gin.Context, dest string, file *multipart.FileHeader) string {

	fullPath := dest + "/" + fmt.Sprint(time.Now().Unix()) + "/" + file.Filename
	fmt.Println(fullPath)
	AbortError(c.SaveUploadedFile(file, fullPath))
	return fullPath
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
