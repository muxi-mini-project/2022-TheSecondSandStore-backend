package goods

import (
	"second/model"
)

func Search(UserId int, GoodsId int) bool {

	var collections []model.Collection
	if err := model.MysqlDb.Db.Where("owner_id = ?", UserId).Where("goods_id = ?", GoodsId).Find(&collections).Error; err != nil {
		return false
	}

	return true
}

func SearchAllGoods(userid int, super model.Super) {
	model.MysqlDb.Db.Where("seller_id = ?", userid).Order("id DESC").First(super)
}

func Save(goodsid int, super model.Super) {
	model.MysqlDb.Db.Where("id = ?", goodsid).Save(super)
}

func SearchWithCondition(constr string, super interface{}) {
	model.MysqlDb.Db.Where("description LIKE ?", constr).Where("if_del = ?", false).Where("if_sell = ?", false).Find(super)
}
