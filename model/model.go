package model

type User struct {
	UserId    int    `json:"user_id" gorm:"user_id"`
	Account   string `json:"account" gorm:"account"`
	NickName  string `json:"nick_name" gorm:"nick_name"`
	Image     string `json:"image" gorm:"image"`
	QQAccount string `json:"qq_account" gorm:"qq_account"`
}

type Goods struct {
	GoodsId     int    `json:"goods_id" gorm:"goods_id"`
	BuyerId     int    `json:"buyer_id" gorm:"buyer_id"`
	SellerId    int    `json:"seller_id" gorm:"seller_id"`
	Time        string `json:"time" gorm:"time"`
	Description string `json:"description" gorm:"description"`
	Image       string `json:"image" gorm:"image"`
	TagIds      string `json:"tag_ids" gorm:"tag_ids"`
}

type Tag struct {
	TagId   int    `json:"tag_id" gorm:"tag_id"`
	OwnerId string `json:"owner_id" gorm:"owner_id"`
	Content string `json:"content" gorm:"content"`
}

type Collection struct {
	CollectionId int `json:"collection_id" gorm:"collection_id"`
	OwnerId      int `json:"owner_id" gorm:"owner_id"`
	GoodsId      int `json:"goods_id" gorm:"goods_id"`
}

type Feedback struct {
	FeedbackId int    `json:"feedback_id" gorm:"feedback_id"`
	OwnerId    int    `json:"owner_id" gorm:"owner_id"`
	Content    string `json:"content" gorm:"content"`
}
