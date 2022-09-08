package config

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB connect
//var (
//	db *gorm.DB
//)
//
//func Connect() {
//	d, err := gorm.Open("mysql", "root:password@tcp(localhost:6603)/test?charset=utf8&parseTime=true&loc=Local")
//	if err != nil {
//		panic(err)
//	}
//
//	db = d
//}
//
//func GetDB() *gorm.DB {
//	return db
//}
