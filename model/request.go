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
	TagIds      []int    `json:"tag_ids"`
	Time        string   `json:"time"`
}

type TagInfo struct {
	Content string `json:"content"`
}

type UserInfo struct {
	Image    string `json:"image"`
	Nickname string `json:"nickname"`
}
