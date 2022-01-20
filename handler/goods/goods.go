package goods

import (
	"log"
	"second/handler"
	"second/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsInfo struct {
	BuyerId string `json:"buyer_id"`
}

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
		log.Fatal(err)
	}

	info := model.GoodsInfo{}
	err = c.BindJSON(&info)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
		log.Fatal(err)
	}

	goods := model.Goods{}

	if err := model.MysqlDb.Db.Where("id = ?", id).First(&goods).Error; err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Because of some errors,it has failed to be found",
			Data:    "null",
		})
		log.Fatal(err)
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
		log.Fatal(err)
	}

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "successful",
	})

}

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
		log.Fatal(err)
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
		res.UserNickname = superuser.NickName

		Res = append(Res, res)
	}

	str, err := handler.ObjectToString(Res)

	if err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "some errors happened in the server",
			Data:    "null",
		})
		log.Fatal(err)
	}

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    str,
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

}

// @Summary 发布商品
// @Description 发布一个新商品
// @Tags goods
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body goods.GoodsInfo true "GoodsInfo"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /goods [post]
func NewOne(c *gin.Context) {
	info := model.GoodsInfo{}
	err := c.BindJSON(&info)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
	}

	useridstr := c.Request.Header.Get("userID")
	userid, err := strconv.Atoi(useridstr)
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    "null",
		})
		log.Fatal(err)
	}

	goods := model.Goods{}
	goods.SellerId = userid
	goods.Image = info.Image
	goods.TagIds = info.TagIds
	goods.Time = info.Time
	goods.Description = info.Description
	if err := model.MysqlDb.Db.Create(&goods).Error; err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Because of some errors,it has failed to be created",
			Data:    "null",
		})
		log.Fatal(err)
	}

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
// @Success 200 "ok,it has been searched and provided successfully"
// @Failure 400 "errors in server"
// @Router /goods/details/all/condition/:condition [get]
func GetInfoCond(c *gin.Context) {

}