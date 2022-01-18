package router

import (
	"net/http"
	"second/handler/user"
	"second/router/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func LoadRouters(r *gin.Engine) {

	//加载中间件和失败路由

	r.Use(gin.Recovery(), middleware.NoCache)
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//swagger

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//路由

	authrouter := r.Group("/api/v1/auth")
	{
		authrouter.POST("/login", user.Login)
	}
}
