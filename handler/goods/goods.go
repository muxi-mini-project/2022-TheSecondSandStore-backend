package goods

import (
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
// @Param info body goods.GoodsInfo true "GoodsInfo"
// @Param goods_id path int true "商品ID"
// @Success 200 "ok,it has been updated successfully"
// @Failure 400 "errors in server"
// @Router /goods/details/one/:goods_id [put]
func UpdateInfo(c *gin.Context) {

}

// @Summary 获取信息
// @Description 获取所有商品信息
// @Tags goods
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "ok,it has been provided successfully"
// @Failure 400 "errors in server"
// @Router /goods/details/all [get]
func GetInfoAll(c *gin.Context) {

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
// @Success 200 "ok,it has been added successfully"
// @Failure 400 "errors in server"
// @Router /goods [post]
func NewOne(c *gin.Context) {

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
