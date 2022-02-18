package model

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CollectionResponse struct {
	UserImage         string   `json:"user_image"`
	UserNickname      string   `json:"user_nickname"`
	QQAccount         string   `json:"qq_account"`
	Time              string   `json:"time"`
	Content           string   `json:"content"`
	GoodsImagesVideos []string `json:"goods_images_videos"`
	IfSell            bool     `json:"if_sell"`
	IfDel             bool     `json:"if_del"`
}

type GoodsResponse struct {
	UserImage         string   `json:"user_image"`
	UserNickname      string   `json:"user_nickname"`
	QQAccount         string   `json:"qq_account"`
	Time              string   `json:"time"`
	Content           string   `json:"content"`
	GoodsImagesVideos []string `json:"goods_images_videos"`
	IfSell            bool     `json:"if_sell"`
	IfDel             bool     `json:"if_del"`
	IfCollected       bool     `json:"if_collected"`
	Price             string   `json:"price"`
}

type TagResponse struct {
	TagId   int    `json:"tag_id"`
	Content string `json:"content"`
}

type UserResponse struct {
	Image             string `json:"image"`
	Nickname          string `json:"nickname"`
	CollectionsNumber int    `json:"collections_number"`
	PostsNumber       int    `json:"posts_number"`
	SellsNumber       int    `json:"sells_number"`
}
