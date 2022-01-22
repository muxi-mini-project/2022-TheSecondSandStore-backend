package token

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	maxage = 60 * 60 * 24 * 7
)

var jwtKey []byte

// 创建将被编码为JWT的结构。
// 我们将jwt.StandardClaims作为嵌入式类型，以提供到期时间等字段。
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(c *gin.Context, userid int, account string) error {

	jwtKey = []byte(viper.GetString("token.secret_key"))
	// 在这里声明令牌的到期时间，我们将其保留为一周
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	// 创建JWT声明，其中包括用户名和有效时间

	claims := &Claims{
		Username: account,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			Id:        strconv.Itoa(userid),
			Issuer:    "rosy",
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 使用用于签名的算法和令牌

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 创建JWT字符串

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

		// 如果创建JWT时出错，则返回内部服务器错误

		return err
	}

	// 最后，我们将客户端cookie token设置为刚刚生成的JWT
	// 我们还设置了与令牌本身相同的cookie到期时间
	host := viper.GetString("web.host")
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Domain:   host,
		Path:     "/",
	})

	return nil
}

func Parse(c *gin.Context) (*Claims, error) {

	// 初始化`Claims`实例
	claims := &Claims{}

	// 我们可以从每个请求的Cookie中获取会话令牌
	value, err := c.Cookie("token")
	if err != nil {
		return claims, err
	}

	// 从Cookie获取JWT字符串
	tknStr := value

	// 解析JWT字符串并将结果存储在`claims`中。
	// 请注意，我们也在此方法中传递了密钥。
	// 如果令牌无效（如果令牌已根据我们设置的登录到期时间过期）或者签名不匹配,此方法会返回错误.
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return claims, errors.New("This token has been expired.")
	}
	if !tkn.Valid {
		return claims, errors.New("the token is invalid")
	}

	return claims, nil
}
