package collection

import (
	"log"
	_ "second/handler"
	"second/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 添加收藏
// @Description 将创建的收藏加入数据库中
// @Tags collection
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body model.CollectionInfo true "The id of the goods"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /collection [post]
func CreateCollection(c *gin.Context) {
	info := model.CollectionInfo{}
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
	goodsid := info.GoodsId

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

	collection := model.Collection{}
	collection.GoodsId = goodsid
	collection.OwnerId = userid
	if err := model.Create(&collection); err != nil {
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

// @Summary 删除收藏
// @Description 将收藏从数据库中删除
// @Tags collection
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param collection_id path string true "用户在搜索框内输入的搜索内容"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /collection/{collection_id} [delete]
func DeleteOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("collection_id"))
	if err != nil {
		c.JSON(400, model.Response{
			Code:    400,
			Message: "some errors in the body of the request",
			Data:    nil,
		})
		log.Println(err)
		return
	}
	collection := model.Collection{}
	if err := model.Delete(id, "id", &collection); err != nil {
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

// @Summary 获得收藏
// @Description 获取用户所有收藏
// @Tags collection
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /collection [get]
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

	var collections []model.Collection

	if err := model.Find(userid, "owner_id", &collections); err != nil {
		c.JSON(200, model.Response{
			Code:    200,
			Message: "ok,empty",
			Data:    nil,
		})
		return
	}

	var Res []model.CollectionResponse
	superuser := &model.SuperUser
	superuser.AutoUpdate(userid)

	for _, v := range collections {
		var res = model.CollectionResponse{}

		goodsid := v.GoodsId

		super := &model.SuperGoods
		super.AutoUpdate(goodsid)

		res.Content = super.Description
		res.GoodsImagesVideos = StringToStringSlice(super.ImagesVideos)
		res.Time = super.Time
		res.IfDel = super.IfDel
		res.IfSell = super.IfSell

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
