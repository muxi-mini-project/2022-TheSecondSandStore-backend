package model

type User struct {
	Id        int    `json:"id" gorm:"column:id"`
	Account   string `json:"account" gorm:"column:account"`
	Nickname  string `json:"nickname" gorm:"column:nickname"`
	Password  string `json:"password" gorm:"column:password"`
	Image     string `json:"image" gorm:"column:image"`
	QQAccount string `json:"qq_account" gorm:"column:qq_account"`
}

type Goods struct {
	Id           int    `json:"id" gorm:"column:id"`
	BuyerId      int    `json:"buyer_id" gorm:"column:buyer_id"`
	SellerId     int    `json:"seller_id" gorm:"column:seller_id"`
	Time         string `json:"time" gorm:"column:time"`
	Description  string `json:"description" gorm:"column:description"`
	ImagesVideos string `json:"images_videos" gorm:"column:images_videos"`
	TagIds       string `json:"tag_ids" gorm:"column:tag_ids"`
	IfSell       bool   `json:"if_sell" gorm:"column:if_sell"`
	IfDel        bool   `json:"if_del" gorm:"column:if_del"`
	Price        string `json:"price" gorm:"column:price"`
}

type Tag struct {
	Id      int    `json:"id" gorm:"column:id"`
	OwnerId int    `json:"owner_id" gorm:"column:owner_id"`
	Content string `json:"content" gorm:"column:content"`
}

type Collection struct {
	Id      int `json:"id" gorm:"column:id"`
	OwnerId int `json:"owner_id" gorm:"column:owner_id"`
	GoodsId int `json:"goods_id" gorm:"column:goods_id"`
}

type Feedback struct {
	Id      int    `json:"id" gorm:"column:id"`
	OwnerId int    `json:"owner_id" gorm:"column:owner_id"`
	Content string `json:"content" gorm:"column:content"`
}
