package middleware

import (
	"second/model"
	"second/pkg/token"
	"strconv"
	"log"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Parse the json web token.
	//ctx, err := auth.ParseRequest(c)
	ctx, err := token.Parse(c)
	if err != nil {
		c.JSON(401, model.Response{
			Code:    401,
			Message: err.Error(),
			Data:    "null",
		})
		log.Println(err)
		c.Abort()
		return
	}
	id, err := strconv.Atoi(ctx.Id)
	if err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "errors in the server",
			Data:    "null",
		})
		c.Abort()
		log.Println(err)
		return
	}
	c.Set("userID", id)

	c.Next()
}
