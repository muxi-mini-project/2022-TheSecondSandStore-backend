package model

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrorNotFound = gorm.ErrRecordNotFound
)

func Create(super Super) error {
	err := MysqlDb.Db.Create(super).Error
	return err
}

func Query(value interface{}, key string, super Super) error {
	query := fmt.Sprintf("%s = ?", key)
	err := MysqlDb.Db.Where(query, value).First(super).Error
	return err
}

func Save(value interface{}, key string, super Super) error {
	query := fmt.Sprintf("%s = ?", key)
	err := MysqlDb.Db.Where(query, value).Save(super).Error
	return err
}

func Delete(value interface{}, key string, super Super) error {
	query := fmt.Sprintf("%s = ?", key)
	err := MysqlDb.Db.Where(query, value).Delete(super).Error
	return err
}

func Find(value interface{}, key string, super interface{}) error {
	query := fmt.Sprintf("%s = ?", key)
	err := MysqlDb.Db.Where(query, value).Find(super).Error
	return err
}

func Search(super interface{}) error {
	err := MysqlDb.Db.Order("id DESC").Where("if_del = ?", false).Where("if_sell = ?", false).Find(super).Error
	return err
}

func First(value interface{}, key string, super Super) error {
	query := fmt.Sprintf("%s = ?", key)
	err := MysqlDb.Db.Where(query, value).First(super).Error
	return err
}
