package feedback

import (
	"log"
	"second/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 添加反馈
// @Description 将创建的意见反馈加入数据库中
// @Tags feedback
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body model.FeedbackInfo true "FeedbackInfo"
// @Success 200 {object} model.Response "successful"
// @Failure 400 {object} model.Response "errors!"
// @Failure 401 {object} model.Response "Errors in authentication by token"
// @Failure 500 {object} model.Response "errors!"
// @Router /feedback [post]
func CreateFeedback(c *gin.Context) {
	info := model.FeedbackInfo{}
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

	content := info.Content

	feedback := model.Feedback{}
	feedback.Content = content
	feedback.OwnerId = userid
	if err := model.MysqlDb.Db.Create(&feedback).Error; err != nil {
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
