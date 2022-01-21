package token

import (
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var (
	JWTKEY           string
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
	ErrTokenInvalid  = errors.New("the token is invalid")
	ErrTokenExpired  = errors.New("the token is expired")
)

type Claims struct {
	UserId int
	jwt.StandardClaims
}

type TokenPayload struct {
	ID      int           `json:"id"`
	Expired time.Duration `json:"expired"` // 有效时间（nanosecond）
}

type TokenClaims struct {
	ID        int   `json:"id"`
	ExpiresAt int64 `json:"expires_at"`
}

type TokenResolve struct {
	ID        int   `json:"id"`
	ExpiresAt int64 `json:"expires_at"` // 过期时间（时间戳，10位）
}

func GenerateToken(userid int) (string, error) {
	key := viper.Get("token.secret_key")
	JWTKEY = key.(string)
	jwtkey := []byte(JWTKEY)
	payload := &TokenPayload{
		ID:      userid,
		Expired: GetExpiredTime(),
	}
	claims := &TokenClaims{
		ID:        payload.ID,
		ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tokenString, nil
}

/*func ParseToken(token string) (*Claims, error) { SigningMethodHS256
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKEY), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}*/

func ResolveToken(tokenStr string) (*TokenResolve, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTKEY), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	t := &TokenResolve{
		ID:        claims.ID,
		ExpiresAt: claims.ExpiresAt,
	}
	return t, nil
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

func GetExpiredTime() time.Duration {
	day := viper.GetInt("token.expired")
	return time.Hour * 24 * time.Duration(day)
}
