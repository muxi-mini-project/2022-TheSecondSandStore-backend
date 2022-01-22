package user

import (
	"encoding/base64"
	"strconv"

	"github.com/spf13/viper"
)

func DefaultImage() string {
	port := viper.GetInt("web.port")
	portstr := strconv.Itoa(port)
	host := viper.GetString("web.host")
	return host + ":" + portstr + "/static/users/default.jpg"
}

func DefaultNickname(id int) string {
	sec := base64.StdEncoding.EncodeToString([]byte("Second"))
	str := "用户" + sec + strconv.Itoa(id)
	return str
}
