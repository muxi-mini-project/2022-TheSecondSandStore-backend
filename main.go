package main

import (
	"second/config"
	"second/model"
	"second/router"

	"github.com/gin-gonic/gin"
)

// @title TheSecondSandStore
// @version 1.0.0
// @description 店小贰小程序
// @termsOfService http://swagger.io/terrms/
// @contact.name rosy
// @contact.email 2313661940@qq.com
// @BasePath /api/v1
// @Host 121.41.45.139:80
// @Schemes http
func main() {
	config.ConfigInit()
	model.InitConnection()
	r := gin.Default()
	router.LoadRouters(r)
	r.Run(":80")
}
