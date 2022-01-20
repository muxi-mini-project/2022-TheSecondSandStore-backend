package middleware

import (
	"second/model"
	"second/pkg/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Parse the json web token.
	ctx, err := auth.ParseRequest(c)
	if err != nil {
		c.JSON(401, model.Response{
			Code:    401,
			Message: "Errors in authentication by token",
			Data:    "null",
		})
		c.Abort()
		return
	}
	c.Set("userID", ctx.ID)
	c.Set("expiresAt", ctx.ExpiresAt)

	c.Next()
}
