package model

import (
	"gorm.io/gorm"
)

type SqlDb struct {
	Db *gorm.DB
}

var MysqlDb *SqlDb

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
}
