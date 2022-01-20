package model

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type CollectionInfo struct {
	GoodsId int `json:"goods_id"`
}

type FeedbackInfo struct {
	Content string
}

type GoodsInfo struct {
	Description string
	Image       string
	TagIds      string
	Time        string
}
