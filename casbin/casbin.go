package casbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

var E *casbin.Enforcer

func init() {
	// 连接数据库
	db, err := gormadapter.NewAdapter("sqlite3", "./gorm.db")
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}
	e, err := casbin.NewEnforcer("../config/model.conf", db)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		return
	}
	E = e
}
