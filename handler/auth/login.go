package auth

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// @Summary 输入账号密码登录
// @Description 登录
// @Tags auth
// @Accept  json
// @Produce  json
// @Param payload body auth.LoginRequest true "Account 账户 Password 密码"
// @Success 200 "login successfully"
// @Failure 400 "errors in request"
// @Failure 400 "not found"
// @Failure 500 "errors in server"
// @Router /auth [post]
func Login(c *gin.Context) {

}
