package main

import (
	_ "fmt"
	"second/config"
	"second/model"

	"github.com/gin-gonic/gin"
	_ "github.com/spf13/viper"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title TheSecondSandStore
// @version 1.0.0
// @description 店小贰小程序
// @termsOfService http://swagger.io/terrms
// @contact.name rosy
// @contact.email 2313661940@qq.com
// @host localhost:8080
// @BasePath:/api/v1
// @Schemes http

func main() {
	config.ConfigInit()
	model.InitConnection()
	r := gin.Default()
	r.Run(":8080")
}
