package auth

import (
	"encoding/base64"
	_ "fmt"
	"log"
	_ "second/handler"
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
// @Param req body model.LoginRequest true "Account 账户 Password 密码"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors"
// @Failure 401 {object} model.Response "Unauthentication"
// @Failure 500 {object} model.Response "errors!"
// @Router /auth [post]
func Login(c *gin.Context) {
	req := model.LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "errors in the server",
			Data:    nil,
		})
		log.Println(err)
		return
	}
	if req.Account == "" || req.Password == "" {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "errors!Account or password is empty.",
			Data:    nil,
		})
		return
	}

	var user = model.User{}

	if err := model.Query(req.Account, "account", &user); err != nil {
		if err == model.ErrorNotFound {

			pwd := base64.StdEncoding.EncodeToString([]byte(req.Password))

			_, err = ccnu.GetUserInfoFormOne(req.Account, req.Password)
			if err != nil {
				c.JSON(400, model.Response{
					Code:    400,
					Message: "The password or the account is wrong",
					Data:    nil,
				})
				log.Println(err)
				return
			}
			user.Account = req.Account
			user.Password = pwd
			if err := model.Create(&user); err != nil {

				c.JSON(500, model.Response{
					Code:    500,
					Message: "some errors in the server",
					Data:    nil,
				})
				log.Println(err)
				return
			}
			if err := model.Query(req.Account, "account", &user); err != nil {
				c.JSON(500, model.Response{
					Code:    500,
					Message: "some errors in the server",
					Data:    nil,
				})
				log.Println(err)
				return
			}

			if err := Init(user.Id, req.Account); err != nil {
				c.JSON(500, model.Response{
					Code:    500,
					Message: "some errors in the server",
					Data:    nil,
				})
				log.Println(err)
				return
			}
		} else {
			c.JSON(500, model.Response{
				Code:    500,
				Message: "some errors in the server",
				Data:    nil,
			})
			log.Println(err)
			return
		}
	} else {
		if user.Password != base64.StdEncoding.EncodeToString([]byte(req.Password)) {
			c.JSON(400, model.Response{
				Code:    400,
				Message: "The password or the account is wrong",
				Data:    nil,
			})
			log.Println(err)
			return
		}
	}

	if err := model.Query(req.Account, "account", &user); err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "some errors in the server",
			Data:    nil,
		})
		log.Println(err)
		return
	}

	err := token.Signin(c, user.Id, req.Account)

	if err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Failed to generate token",
			Data:    nil,
		})
		log.Println(err)
		return
	}

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "successful",
	})

}
