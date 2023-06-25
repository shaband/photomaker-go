package helpers

import (
	"encoding/json"
	"fmt"
)

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
