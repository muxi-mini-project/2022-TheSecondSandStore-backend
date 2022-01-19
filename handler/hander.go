package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func SendRes(c *gin.Context, code string, message string, data string) {
	str, _ := strconv.Atoi(code)
	c.JSON(str, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
