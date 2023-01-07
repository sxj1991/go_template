package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"log"
)

func main() {
	// 连接数据库
	db, err := adapter.NewAdapter("sqlite3", "./gorm.db")
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}
	e, err := casbin.NewEnforcer("./casbin/v2/model.conf", db)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		return
	}
	//从DB加载策略
	e.LoadPolicy()

	// 更新策略
	e.AddPolicy("admin", "data", "read")
	e.AddRoleForUser("zhangsan", "admin")
	// e.RemovePolicy(...)

	// 保存到数据库
	e.SavePolicy()

	//判断策略中是否存在
	// 请求
	sub := "zhangsan" // 想要访问资源的用户。
	obj := "data"     // 将被访问的资源。
	act := "write"    // 用户对资源执行的操作。
	if ok, _ := e.Enforce(sub, obj, act); ok {
		fmt.Println("恭喜您,权限验证通过")
	} else {
		fmt.Println("很遗憾,权限验证没有通过")
	}
}
