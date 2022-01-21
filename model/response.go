package model

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

type TagResponse struct {
	TagId   int    `json:"tag_id"`
	Content string `json:"content"`
}

type UserResponse struct {
	Image    string `json:"image"`
	Nickname string `json:"nickname"`
}
