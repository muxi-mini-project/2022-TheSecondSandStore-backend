package middleware

import (
	"second/handler"
	"second/pkg/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Parse the json web token.
	ctx, err := auth.ParseRequest(c)
	if err != nil {
		handler.SendRes(c, "401", "Errors in authentication by token", "")
		c.Abort()
		return
	}

	c.Set("userID", ctx.ID)
	c.Set("expiresAt", ctx.ExpiresAt)

	c.Next()
}
