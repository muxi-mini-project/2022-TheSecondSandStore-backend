package model

type User struct {
	Id        int    `json:"id" gorm:"id"`
	Account   string `json:"account" gorm:"account"`
	NickName  string `json:"nick_name" gorm:"nick_name"`
	Password  string `json:"password" gorm:"password"`
	Image     string `json:"image" gorm:"image"`
	QQAccount string `json:"qq_account" gorm:"qq_account"`
}

type Goods struct {
	Id          int    `json:"id" gorm:"id"`
	BuyerId     int    `json:"buyer_id" gorm:"buyer_id"`
	SellerId    int    `json:"seller_id" gorm:"seller_id"`
	Time        string `json:"time" gorm:"time"`
	Description string `json:"description" gorm:"description"`
	Image       string `json:"image" gorm:"image"`
	TagIds      string `json:"tag_ids" gorm:"tag_ids"`
}

type Tag struct {
	Id      int    `json:"id" gorm:"id"`
	OwnerId string `json:"owner_id" gorm:"owner_id"`
	Content string `json:"content" gorm:"content"`
}

type Collection struct {
	Id      int `json:"id" gorm:"id"`
	OwnerId int `json:"owner_id" gorm:"owner_id"`
	GoodsId int `json:"goods_id" gorm:"goods_id"`
}

type Feedback struct {
	Id      int    `json:"id" gorm:"id"`
	OwnerId int    `json:"owner_id" gorm:"owner_id"`
	Content string `json:"content" gorm:"content"`
}
