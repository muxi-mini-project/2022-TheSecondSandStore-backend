package auth

import (
	_ "errors"

	_ "second/pkg/token"

	_ "github.com/gin-gonic/gin"
)

/*
var (
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")

	ErrTokenInvalid = errors.New("The token is invalid.")
)

// Context is the context of the JSON web token.
type Context struct {
	ID        int
	ExpiresAt int64 // 过期时间（时间戳，10位）
}

// Parse parses the token, and returns the context if the token is valid.
func Parse(tokenString string) (*Context, error) {
	t, err := token.ResolveToken(tokenString)
	if err != nil {
		return nil, err
	}

	return &Context{
		ID:        t.ID,
		ExpiresAt: t.ExpiresAt,
	}, nil

}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("token")
	if len(header) == 0 {
		return nil, ErrMissingHeader
	}

	return Parse(header)
}
*/
