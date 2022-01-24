package goods

import (
	"second/model"
)

func Search(UserId int, GoodsId int) bool {

	var collections []model.Collection
	if err := model.MysqlDb.Db.Where("owner_id = ?", UserId).Where("goods_id = ?", GoodsId).Find(&collections).Error; err != nil {
		return false
	} else {
		return true
	}

}
