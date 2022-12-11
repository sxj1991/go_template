package database

import (
	"go_template/setting"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		setting.Log.Panicf("初始化数据库报错：%v\n", err)
	}

	DB = db
}
