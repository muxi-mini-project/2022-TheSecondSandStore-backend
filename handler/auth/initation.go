package auth

import (
	"encoding/base64"
	_ "fmt"
	"second/model"
	"strconv"

	"github.com/spf13/viper"
)

func Init(userid int, account string) error {

	//标签初始化

	tag := model.Tag{}
	tag.Content = "#计算机教材"
	tag.OwnerId = userid
	if err := model.Create(&tag); err != nil {
		return err
	}
	tag = model.Tag{}
	tag.Content = "#日用品"
	tag.OwnerId = userid
	if err := model.Create(&tag); err != nil {
		return err
	}
	tag = model.Tag{}
	tag.Content = "#文具"
	tag.OwnerId = userid
	if err := model.Create(&tag); err != nil {
		return err
	}

	//头像以及昵称初始化

	path := "/static/users/default.jpg"
	port := viper.GetInt("web.port")
	portstr := strconv.Itoa(port)
	host := viper.GetString("web.host")
	r := host + ":" + portstr + path
	user := model.User{}
	if err := model.Query(account, "account", &user); err != nil {
		return err
	}
	user.Image = r

	sec := base64.StdEncoding.EncodeToString([]byte("Second"))
	str := "用户" + sec + strconv.Itoa(userid)
	user.Nickname = str

	if err := model.Save(account, "account", &user); err != nil {
		return err
	}
	return nil
}
