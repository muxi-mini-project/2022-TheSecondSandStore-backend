package user

import (
	"second/model"
)

func CountNumbers(userid int) (int, int, int) {
	var coll []model.Collection
	model.MysqlDb.Db.Where("owner_id = ?", userid).Find(&coll)

	var post []model.Goods
	model.MysqlDb.Db.Where("seller_id = ?", userid).Find(&post)

	var sell []model.Goods
	model.MysqlDb.Db.Where("seller_id = ?", userid).Where("if_sell = ?", true).Find(&sell)

	return len(coll), len(post), len(sell)
}
