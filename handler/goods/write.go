package goods

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"
	_ "strings"

	"github.com/spf13/viper"
)

func Write(EncodeStr string, GoodsId int, id int, Type string) (string, error) {
	Bytes, err := base64.StdEncoding.DecodeString(EncodeStr)
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("/static/goods/%d.%d.%s", GoodsId, id, Type)

	ioutil.WriteFile("."+path, Bytes, 0777)

	port := viper.GetInt("web.port")
	PortStr := strconv.Itoa(port)
	host := viper.GetString("web.host")
	return host + ":" + PortStr + path, nil
}
