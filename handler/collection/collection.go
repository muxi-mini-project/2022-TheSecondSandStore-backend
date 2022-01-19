package collection

import (
	"github.com/gin-gonic/gin"
)

type CollectionInfo struct {
	GoodsId      int `json:"goods_id"`
	CollectionId int `json:"collection_id"`
}

// @Summary 添加收藏
// @Description 将创建的收藏加入数据库中
// @Tags collection
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body collection.CollectionInfo true "CollectionInfo"
// @Success 200 "ok,it has been added successfully"
// @Failure 400 "errors in server"
// @Router /collection [post]
func NewOne(c *gin.Context) {

}

// @Summary 删除收藏
// @Description 将收藏从数据库中删除
// @Tags collection
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body collection.CollectionInfo true "CollectionInfo"
// @Success 200 "ok,it has been deleted successfully"
// @Failure 400 "errors in server"
// @Router /collection [delete]
func DeleteOne(c *gin.Context) {

}

// @Summary 获得收藏
// @Description 获取用户所有收藏
// @Tags collection
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "ok,it has been provided successfully"
// @Failure 400 "errors in server"
// @Router /collection [get]
func GetInfo(c *gin.Context) {

}
