package goods

import (
	"log"
	_ "second/handler"
	"second/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsInfo struct {
	BuyerId string `json:"buyer_id"`
}

/*
// @Summary 更新信息
// @Description 修改商品信息
// @Tags goods
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body model.GoodsInfo true "GoodsInfo"
// @Param goods_id path int true "商品ID"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /goods/details/one/:goods_id [put]
func UpdateInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
		log.Println(err)
		return
	}

	info := model.GoodsInfo{}
	err = c.BindJSON(&info)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
		log.Println(err)
		return
	}

	goods := model.Goods{}

	if err := model.MysqlDb.Db.Where("id = ?", id).First(&goods).Error; err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Because of some errors,it has failed to be found",
			Data:    "null",
		})
		log.Println(err)
		return
	}

	goods.Image = info.Image
	goods.TagIds = info.TagIds
	goods.Time = info.Time
	goods.Description = info.Description

	if err := model.MysqlDb.Db.Where("id = ?", id).Save(&goods).Error; err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Because of some errors,it has failed to be saved",
			Data:    "null",
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
*/
// @Summary 获取信息
// @Description 获取所有商品信息
// @Tags goods
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /goods/details/all [get]
func GetInfoAll(c *gin.Context) {
	var goodses []model.Goods

	if err := model.MysqlDb.Db.Order("id DESC").Find(&goodses).Error; err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Because of some errors,it has failed to be deleted",
			Data:    "null",
		})
		log.Println(err)
		return
	}

	var Res []model.GoodsResponse

	for _, v := range goodses {
		var res = model.GoodsResponse{}

		res.Content = v.Description
		res.GoodsImage = v.Image
		res.Time = v.Time

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

// @Summary 获取信息
// @Description 获取某一商品信息
// @Tags goods
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param goods_id path int true "商品ID"
// @Success 200 "ok,it has been provided successfully"
// @Failure 400 "errors in server"
// @Router /goods/details/one/:goods_id [get]
func GetInfoId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("goods_id"))
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
		log.Println(err)
		return
	}

	supergoods := &model.SuperGoods
	supergoods.AutoUpdate(id)

	res := model.GoodsResponse{}
	res.Content = supergoods.Description
	res.GoodsImage = supergoods.Image
	res.Time = supergoods.Time

	superuser := &model.SuperUser
	superuser.AutoUpdate(supergoods.SellerId)

	res.QQAccount = superuser.QQAccount
	res.UserImage = superuser.Image
	res.UserNickname = superuser.Nickname

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    res,
	})
}

// @Summary 发布商品
// @Description 发布一个新商品
// @Tags goods
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body model.GoodsInfo true "GoodsInfo"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /goods [post]
func CreateGoods(c *gin.Context) {
	info := model.GoodsInfo{}
	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
		log.Println(err)
		return
	}

	UserId, ok := c.Get("userID")
	if !ok {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "errors in the server",
			Data:    "null",
		})
		return
	}
	userid := UserId.(int)

	goods := model.Goods{}
	goods.SellerId = userid
	goods.TagIds = info.TagIds
	goods.Time = info.Time
	goods.Description = info.Description
	if err := model.MysqlDb.Db.Create(&goods).Error; err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Because of some errors,it has failed to be created",
			Data:    "null",
		})
		log.Println(err)
		return
	}

	goods = model.Goods{}
	model.MysqlDb.Db.Where("seller_id = ?", userid).Order("id DESC").First(&goods)

	src, err := Write(info.Image, goods.Id)

	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "errors in the image",
			Data:    "null",
		})
		log.Println(err)
		return
	}

	goods.Image = src
	model.MysqlDb.Db.Where("id = ?", goods.Id).Save(&goods)

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "created successfully",
	})
}

// @Summary 搜索商品
// @Description 获取搜索的商品信息
// @Tags goods
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param  condition path string true "用户在搜索框内输入的搜索内容"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /goods/details/all/condition/:condition [get]
func GetInfoCond(c *gin.Context) {
	condition := c.Param("condition")

	var goods []model.Goods
	constr := "%" + condition + "%"
	model.MysqlDb.Db.Where("description LIKE ?", constr).Find(&goods)

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    goods,
	})
}
