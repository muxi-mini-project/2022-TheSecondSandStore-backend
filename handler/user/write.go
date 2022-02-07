package user

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func Write(EncodeStr string, userid int) (string, error) {
	imageBytes, err := base64.StdEncoding.DecodeString(EncodeStr)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("/static/users/%d.jpg", userid)

	ioutil.WriteFile("."+path, imageBytes, 0777)

	port := viper.GetInt("web.port")
	portstr := strconv.Itoa(port)
	host := viper.GetString("web.host")
	return host + ":" + portstr + path, nil
}

func StringToStringSlice(str string) []string {
	return strings.Split(str, " ")
}
