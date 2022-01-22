package auth

import (
	"encoding/base64"
	_ "fmt"
	"second/model"
	"strconv"

	"github.com/spf13/viper"
)

func Init(userid int, account string) {

	//标签初始化

	tag := model.Tag{}
	tag.Content = "#计算机教材"
	tag.OwnerId = userid
	model.MysqlDb.Db.Create(&tag)
	tag = model.Tag{}
	tag.Content = "#日用品"
	tag.OwnerId = userid
	model.MysqlDb.Db.Create(&tag)
	tag = model.Tag{}
	tag.Content = "#文具"
	tag.OwnerId = userid
	model.MysqlDb.Db.Create(&tag)

	//头像以及昵称初始化

	path := "/static/users/default.jpg"
	port := viper.GetInt("web.port")
	portstr := strconv.Itoa(port)
	host := viper.GetString("web.host")
	r := host + ":" + portstr + path
	user := model.User{}
	model.MysqlDb.Db.Where("account = ?", account).First(&user)
	user.Image = r

	sec := base64.StdEncoding.EncodeToString([]byte("Second"))
	str := "用户" + sec + strconv.Itoa(userid)
	user.Nickname = str

	model.MysqlDb.Db.Where("account = ?", account).Save(&user)
}
