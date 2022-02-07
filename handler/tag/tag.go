package tag

import (
	"log"
	_ "second/handler"
	"second/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 添加标签
// @Description 将创建的标签加入数据库中
// @Tags tag
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body model.TagInfo true "TagInfo"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /tag [post]
func CreateTag(c *gin.Context) {
	info := model.TagInfo{}
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

	UserId, ok := c.Get("userID")
	if !ok {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "errors in the server",
			Data:    nil,
		})
		return
	}
	userid := UserId.(int)

	content := info.Content

	tag := model.Tag{}
	tag.Content = content
	tag.OwnerId = userid
	if err := model.Create(&tag); err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Because of some errors,it has failed to be created",
			Data:    nil,
		})
		log.Println(err)
		return
	}

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    "created successfully",
	})
}

// @Summary 删除标签
// @Description 将标签从数据库中删除
// @Tags tag
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param tag_id path int true " 标签ID"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /tag/{tag_id} [delete]
func DeleteOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    nil,
		})
		log.Println(err)
		return
	}

	tag := model.Tag{}
	if err := model.Delete(id, "id", &tag); err != nil {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "Because of some errors,it has failed to be deleted",
			Data:    nil,
		})
		log.Println(err)
		return
	}

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    nil,
	})
}

// @Summary 获得标签
// @Description 获取用户的所有标签
// @Tags tag
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /tag [get]
func GetInfo(c *gin.Context) {
	UserId, ok := c.Get("userID")
	if !ok {
		c.JSON(500, model.Response{
			Code:    500,
			Message: "errors in the server",
			Data:    nil,
		})
		return
	}
	userid := UserId.(int)

	var tags []model.Tag

	if err := model.Find(userid, "owner_id", &tags); err != nil {
		c.JSON(200, model.Response{
			Code:    200,
			Message: "ok,empty",
			Data:    nil,
		})
		return
	}

	var Res []model.TagResponse

	for _, v := range tags {
		var res = model.TagResponse{}

		res.Content = v.Content
		res.TagId = v.Id

		Res = append(Res, res)
	}

	c.JSON(200, model.Response{
		Code:    200,
		Message: "ok",
		Data:    Res,
	})
}
