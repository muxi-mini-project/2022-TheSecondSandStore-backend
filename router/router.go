package router

import (
	"net/http"
	_ "second/docs"
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

	r.Use(middleware.NoCache, middleware.Options, middleware.Secure)
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//swagger

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//加载静态文件

	r.StaticFS("/static", http.Dir("./static"))

	//路由

	authrouter := r.Group("/api/v1/auth")
	{
		authrouter.POST("", auth.Login)
	}

	userrouter := r.Group("/api/v1/user")
	userrouter.Use(middleware.AuthMiddleware)
	{
		userrouter.PUT("/nickname", user.UpdateInfoNickname)
		userrouter.PUT("/image", user.UpdateInfoImage)
		userrouter.GET("", user.GetInfo)
		userrouter.GET("/goods", user.GetGoodsInfo)
		userrouter.DELETE("/goods/:goods_id", user.DelGoods)
		userrouter.PUT("/goods/:goods_id", user.SellGoods)
	}

	goodsrouter := r.Group("/api/v1/goods")
	goodsrouter.Use(middleware.AuthMiddleware)
	{
		//goodsrouter.PUT("/details/one/:goods_id", goods.UpdateInfo)
		goodsrouter.POST("", goods.CreateGoods)
		goodsrouter.GET("/details/all", goods.GetInfoAll)
		goodsrouter.GET("/details/all/condition/:condition", goods.GetInfoCond)
		goodsrouter.GET("/details/one/:goods_id", goods.GetInfoId)

	}

	collectionrouter := r.Group("/api/v1/collection")
	collectionrouter.Use(middleware.AuthMiddleware)
	{
		collectionrouter.POST("", collection.CreateCollection)
		collectionrouter.DELETE("/:collection_id", collection.DeleteOne)
		collectionrouter.GET("", collection.GetInfo)
	}

	tagrouter := r.Group("/api/v1/tag")
	tagrouter.Use(middleware.AuthMiddleware)
	{
		tagrouter.POST("", tag.CreateTag)
		tagrouter.DELETE("/:tag_id", tag.DeleteOne)
		tagrouter.GET("", tag.GetInfo)
	}

	feedbackrouter := r.Group("/api/v1/feedback")
	feedbackrouter.Use(middleware.AuthMiddleware)
	{
		feedbackrouter.POST("", feedback.CreateFeedback)
	}
}
