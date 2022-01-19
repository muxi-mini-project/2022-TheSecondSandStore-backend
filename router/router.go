package router

import (
	"net/http"
	"second/handler/auth"
	"second/handler/collection"
	"second/handler/feedback"
	"second/handler/goods"
	"second/handler/tag"
	"second/handler/user"
	"second/router/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func LoadRouters(r *gin.Engine) {

	//加载中间件和失败路由

	r.Use(middleware.NoCache)
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//swagger

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//路由

	authrouter := r.Group("/api/v1/auth")
	{
		authrouter.POST("/", auth.Login)
	}

	userrouter := r.Group("/api/v1/user")
	userrouter.Use(middleware.AuthMiddleware)
	{
		userrouter.PUT("/nickname", user.UpdateInfoNickname)
		userrouter.PUT("/image", user.UpdateInfoImage)
		userrouter.GET("/", user.GetInfo)
	}

	goodsrouter := r.Group("/api/v1/goods")
	goodsrouter.Use(middleware.AuthMiddleware)
	{
		goodsrouter.PUT("/details/one/:goods_id", goods.UpdateInfo)
		goodsrouter.POST("/", goods.NewOne)
		goodsrouter.GET("/details/all", goods.GetInfoAll)
		goodsrouter.GET("/details/all/condition/:condition", goods.GetInfoCond)
		goodsrouter.GET("/details/one/:goods_id", goods.GetInfoId)

	}

	collectionrouter := r.Group("/api/v1/collection")
	collectionrouter.Use(middleware.AuthMiddleware)
	{
		goodsrouter.POST("/", collection.NewOne)
		goodsrouter.DELETE("/", collection.DeleteOne)
		goodsrouter.GET("/", collection.GetInfo)
	}

	tagrouter := r.Group("/api/v1/tag")
	tagrouter.Use(middleware.AuthMiddleware)
	{
		goodsrouter.POST("/", tag.NewOne)
		goodsrouter.DELETE("/:tag_id", tag.DeleteOne)
		goodsrouter.GET("/", tag.GetInfo)
	}

	feedbackrouter := r.Group("/api/v1/feedback")
	feedbackrouter.Use(middleware.AuthMiddleware)
	{
		goodsrouter.POST("/", feedback.NewOne)
	}
}
