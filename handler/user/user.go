package user

import (
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Nickname string
	Image    string
}

// @Summary 获得用户信息
// @Description 获取用户昵称和头像
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "ok,it has been provided successfully"
// @Failure 400 "errors in server"
// @Router /user [get]
func GetInfo(c *gin.Context) {

}

// @Summary 更新信息
// @Description 修改用户头像
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body user.UserInfo true "UserInfo"
// @Success 200 "ok,it has been updated successfully"
// @Failure 400 "errors in server"
// @Router /user/image [put]
func UpdateInfoImage(c *gin.Context) {

}

// @Summary 更新信息
// @Description 修改用户昵称
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body user.UserInfo true "UserInfo"
// @Success 200 "ok,it has been updated successfully"
// @Failure 400 "errors in server"
// @Router /user/nickname [put]
func UpdateInfoNickname(c *gin.Context) {

}
