package model

import (
	"gorm.io/gorm"
)

type SqlDb struct {
	Db *gorm.DB
}

type Super interface {
	AutoUpdate()
}

var (
	MysqlDb         SqlDb
	SuperUser       User
	SuperGoods      Goods
	SuperTag        Tag
	SuperCollection Collection
	SuperFeedback   Feedback
)

func InitConnection() {
	/*dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", viper.Get("mysql.user"), viper.Get("mysql.password"), viper.Get("mysql.host"), viper.Get("mysql.port"), viper.Get("mysql.db"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqldb := &SqlDb{
		db,
	}
	MysqlDb = sqldb*/
	SuperInit()
}

func SuperInit() {
	SuperUser = User{}
	SuperTag = Tag{}
	SuperGoods = Goods{}
	SuperFeedback = Feedback{}
	SuperCollection = Collection{}
}

func (Super *User) AutoUpdate(id int) {
	user := User{}
	MysqlDb.Db.Where("id = ?", id).First(&user)
	*Super = user
}

func (Super *Tag) AutoUpdate(id int) {
	tag := Tag{}
	MysqlDb.Db.Where("id = ?", id).First(&tag)
	*Super = tag
}

func (Super *Goods) AutoUpdate(id int) {
	goods := Goods{}
	MysqlDb.Db.Where("id = ?", id).First(&goods)
	*Super = goods
}

func (Super *Feedback) AutoUpdate(id int) {
	feedback := Feedback{}
	MysqlDb.Db.Where("id = ?", id).First(&feedback)
	*Super = feedback
}

func (Super *Collection) AutoUpdate(id int) {
	collection := Collection{}
	MysqlDb.Db.Where("id = ?", id).First(&collection)
	*Super = collection
}
