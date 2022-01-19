package tag

import (
	"github.com/gin-gonic/gin"
)

type TagInfo struct {
	OwnerId int
	Content string
}

// @Summary 添加标签
// @Description 将创建的标签加入数据库中
// @Tags tag
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body tag.TagInfo true "TagInfo"
// @Success 200 "ok,it has been added successfully"
// @Failure 400 "errors in server"
// @Router /tag [post]
func NewOne(c *gin.Context) {

}

// @Summary 删除标签
// @Description 将标签从数据库中删除
// @Tags tag
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param tag_id path int true " 标签ID"
// @Success 200 "ok,it has been deleted successfully"
// @Failure 400 "errors in server"
// @Router /tag/:tag_id [delete]
func DeleteOne(c *gin.Context) {

}

// @Summary 获得标签
// @Description 获取用户的所有标签
// @Tags tag
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "ok,it has been provided successfully"
// @Failure 400 "errors in server"
// @Router /tag [get]
func GetInfo(c *gin.Context) {

}
