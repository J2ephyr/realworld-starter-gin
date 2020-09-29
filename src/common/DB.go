package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 全局禁用表名复数
	db.SingularTable(true)
	db.LogMode(true)
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
