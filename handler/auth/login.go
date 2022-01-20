package auth

import (
	"log"
	"second/handler"
	"second/model"
	"second/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/r-rosy/General/ccnu"
)

// @Summary 输入账号密码登录
// @Description 登录
// @Tags auth
// @Accept  json
// @Produce  json
// @Param req body model.LoginRequest true "Account 账户 Password 将密码进行base64编码后的字符串"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "Unauthentication"
// @Failure 500 {object} model.Response "errors!"
// @Router /auth [post]
func Login(c *gin.Context) {
	req := model.LoginRequest{}
	c.BindJSON(&req)
	if req.Account == "" || req.Password == "" {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "errors!Account or password is empty.",
			Data:    "null",
		})
		return
	}

	var user = model.User{}

	if resu := model.MysqlDb.Db.Where("account = ?", req.Account).First(&user); resu.Error != nil {
		_, err := ccnu.GetUserInfoFormOne(req.Account, req.Password)
		if err != nil {
			c.JSON(400, model.Response{
				Code:    400,
				Message: "The password or the account is wrong",
				Data:    "null",
			})
			return
		}
		user.Account = req.Account
		user.Password = req.Password
		model.MysqlDb.Db.Create(&user)
	} else {

		if user.Password != req.Password {
			c.JSON(400, model.Response{
				Code:    400,
				Message: "The password or the account is wrong",
				Data:    "null",
			})
			return
		}
	}

	model.MysqlDb.Db.Where("account = ?", req.Account).First(&user)
	tokenstr := token.GenerateToken(user.Id)
	var result struct {
		token string
	}
	result.token = tokenstr
	str, err := handler.ObjectToString(result)
	if err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "some errors happened in the server",
			Data:    "",
		})
		log.Fatal(err)
	}
	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    str,
	})

}
