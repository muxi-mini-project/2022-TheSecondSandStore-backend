package goods

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"
	_ "strings"

	"github.com/spf13/viper"
)

func Write(EncodeStr string, goodsid int) (string, error) {
	imageBytes, err := base64.StdEncoding.DecodeString(EncodeStr)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("/static/goods/%d.jpg", goodsid)

	ioutil.WriteFile("."+path, imageBytes, 0777)

	port := viper.GetInt("web.port")
	portstr := strconv.Itoa(port)
	host := viper.GetString("web.host")
	return host + ":" + portstr + path, nil
}
