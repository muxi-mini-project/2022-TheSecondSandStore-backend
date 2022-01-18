package token

import (
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	JWTKEY           string
	ErrMissingHeader = errors.New("The length of the `Authorization` header is zero.")
	ErrTokenInvalid  = errors.New("The token is invalid.")
	ErrTokenExpired  = errors.New("The token is expired.")
)

type Claims struct {
	UserId int
	jwt.StandardClaims
}

type TokenClaims struct {
	ID        int   `json:"id"`
	ExpiresAt int64 `json:"expires_at"`
}

func GenerateToken(userid int) string {
	key := viper.Get("token.secret_key")
	JWTKEY = key.(string)
	jwtkey := []byte(JWTKEY)
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "rosy",       // 签名颁发者
			Subject:   "user_token", //签名主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKEY), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (c TokenClaims) Valid() error {
	now := time.Now().Unix()

	if !c.VerifyExpiresAt(now, false) {
		return ErrTokenExpired
	}

	return nil
}

func (c *TokenClaims) VerifyExpiresAt(now int64, required bool) bool {
	if c.ExpiresAt == 0 {
		return !required
	}
	return now <= c.ExpiresAt
}
