package handler

import (
	"encoding/json"
	_ "second/model"

	_ "github.com/gin-gonic/gin"
)

func ObjectToString(object interface{}) (string, error) {
	byte, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}
