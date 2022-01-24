package model

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type CollectionInfo struct {
	GoodsId int `json:"goods_id"`
}

type FeedbackInfo struct {
	Content string `json:"content"`
}

type GoodsInfo struct {
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Videos      []string `json:"videos"`
	TagIds      []int    `json:"tag_ids"`
	Time        string   `json:"time"`
	QQAccount   string   `json:"qq_account"`
}

type TagInfo struct {
	Content string `json:"content"`
}

type UserInfo struct {
	Image    string `json:"image"`
	Nickname string `json:"nickname"`
}
