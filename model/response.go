package model

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type CollectionResponse struct {
	UserImage    string `json:"user_image"`
	UserNickname string `json:"user_nickname"`
	QQAccount    string `json:"qq_account"`
	Time         string `json:"time"`
	Content      string `json:"content"`
	GoodsImage   string `json:"goods_image"`
}

type GoodsResponse struct {
	UserImage    string `json:"user_image"`
	UserNickname string `json:"user_nickname"`
	QQAccount    string `json:"qq_account"`
	Time         string `json:"time"`
	Content      string `json:"content"`
	GoodsImage   string `json:"goods_image"`
}
