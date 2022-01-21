package user

import (
	"log"
	_ "second/handler"
	"second/model"
	"strconv"

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
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /user [get]
func GetInfo(c *gin.Context) {
	UserId, ok := c.Get("userID")
	if !ok {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "errors in the server",
			Data:    "null",
		})
	}
	userid := UserId.(int)

	var res model.UserResponse

	superuser := &model.SuperUser
	superuser.AutoUpdate(userid)

	res.Image = superuser.Image
	res.Nickname = superuser.NickName

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    res,
	})
}

// @Summary 更新信息
// @Description 修改用户头像
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body model.UserInfo true "UserInfo"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /user/image [put]
func UpdateInfoImage(c *gin.Context) {
	info := model.UserInfo{}
	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
	}

	UserIdStr := c.Request.Header.Get("userID")
	userid, err := strconv.Atoi(UserIdStr)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
		log.Fatal(err)
	}

	user := model.User{}
	model.MysqlDb.Db.Where("id = ?", userid).First(&user)
	user.Image = info.Image
	model.MysqlDb.Db.Where("id = ?", userid).Save(&user)

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "successful",
	})
}

// @Summary 更新信息
// @Description 修改用户昵称
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body user.UserInfo true "UserInfo"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /user/nickname [put]
func UpdateInfoNickname(c *gin.Context) {
	info := model.UserInfo{}
	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
	}

	UserIdStr := c.Request.Header.Get("userID")
	userid, err := strconv.Atoi(UserIdStr)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
		log.Fatal(err)
	}

	user := model.User{}
	model.MysqlDb.Db.Where("id = ?", userid).First(&user)
	user.NickName = info.NickName
	model.MysqlDb.Db.Where("id = ?", userid).Save(&user)

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "successful",
	})
}
