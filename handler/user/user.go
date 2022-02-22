package user

import (
	"log"
	_ "second/handler"
	"second/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获得用户信息
// @Description 获取用户昵称头像以及收藏数发布数卖出数
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
	UserId := c.MustGet("userID")
	userid := UserId.(int)

	var res model.UserResponse

	superuser := &model.SuperUser
	superuser.AutoUpdate(userid)
	if superuser.Image == "" {
		res.Image = DefaultImage()
	} else {
		res.Image = superuser.Image
	}

	if superuser.Nickname == "" {
		res.Nickname = DefaultNickname(userid)
	} else {
		res.Nickname = superuser.Nickname
	}

	colle, post, sell := CountNumbers(userid)
	res.CollectionsNumber = colle
	res.PostsNumber = post
	res.SellsNumber = sell
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
			Data:    nil,
		})
		log.Println(err)
		return
	}

	UserId := c.MustGet("userID")
	userid := UserId.(int)

	user := model.User{}
	model.First(userid, "id", &user)
	if info.Image == "" {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "image in the body of your request is lost",
			Data:    nil,
		})
		return
	}

	src, err := Write(info.Image, userid)

	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "errors in the image",
			Data:    nil,
		})
		log.Println(err)
		return
	}

	user.Image = src
	model.Save(userid, "id", &user)
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
// @Param info body model.UserInfo true "UserInfo"
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
			Data:    nil,
		})
	}

	UserId := c.MustGet("userID")
	userid := UserId.(int)

	user := model.User{}
	model.First(userid, "id", &user)
	if info.Nickname == "" {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "the nickname in the body of your request is lost",
			Data:    nil,
		})
		return
	}
	user.Nickname = info.Nickname
	model.Save(userid, "id", &user)
	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "successful",
	})
}

// @Summary 获得用户发布的商品信息
// @Description 获取用户发布的商品信息
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /user/goods [get]
func GetGoodsInfo(c *gin.Context) {
	UserId := c.MustGet("userID")
	userid := UserId.(int)

	var goodses []model.Goods
	if err := model.Find(userid, "seller_id", &goodses); err != nil {
		c.JSON(200, model.Response{
			Code:    200,
			Message: "ok,empty",
			Data:    nil,
		})
		return
	}

	var Res []model.GoodsResponse

	for _, v := range goodses {
		var res = model.GoodsResponse{}

		res.Content = v.Description

		res.GoodsImagesVideos = StringToStringSlice(v.ImagesVideos)
		res.Time = v.Time
		res.IfSell = v.IfSell
		res.IfDel = v.IfDel

		superuser := &model.SuperUser
		superuser.AutoUpdate(v.SellerId)

		res.QQAccount = superuser.QQAccount
		res.UserImage = superuser.Image
		res.UserNickname = superuser.Nickname

		Res = append(Res, res)
	}

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    Res,
	})
}

// @Summary 删除用户发布的商品信息
// @Description 删除用户发布的商品信息
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /user/goods/{goods_id} [delete]
func DelGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    nil,
		})
		log.Println(err)
		return
	}

	super := &model.SuperGoods
	super.AutoUpdate(id)
	super.IfDel = true
	super.Save()

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    nil,
	})
}

// @Summary 用户确认卖出商品
// @Description 用户确认卖出商品
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /user/goods/{goods_id} [put]
func SellGoods(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    nil,
		})
		log.Println(err)
		return
	}

	super := &model.SuperGoods
	super.AutoUpdate(id)
	super.IfSell = true
	super.Save()

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    nil,
	})

}
