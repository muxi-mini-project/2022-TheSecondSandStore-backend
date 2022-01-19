package feedback

import (
	"github.com/gin-gonic/gin"
)

type FeedbackInfo struct {
	Content string `json:"content"`
}

// @Summary 添加反馈
// @Description 将创建的意见反馈加入数据库中
// @Tags feedback
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param info body feedback.FeedbackInfo true "FeedbackInfo"
// @Success 200 "ok,it has been added successfully"
// @Failure 400 "errors in server"
// @Router /feedback [post]
func NewOne(c *gin.Context) {

}
