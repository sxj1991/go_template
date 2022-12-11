package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

type user struct {
	sub string // 想要访问资源的用户
	obj string // 将被访问的资源
	act string // 用户对资源执行的操作
}

func main() {
	cas, err1 := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")

	if err1 != nil {
		fmt.Printf("err:%v\n", err1)
	}

	u1 := user{
		sub: "admin",
		obj: "system:role",
		act: "list",
	}

	ok, err2 := cas.Enforce(u1.sub, u1.obj, u1.act)

	if err2 != nil {
		// 处理err
		fmt.Printf("err:%v\n", err2)
	}

	if ok == true {
		// 允许操作
		fmt.Printf("允许操作")
	} else {
		// 拒绝请求，抛出异常
		fmt.Printf("拒绝请求，抛出异常")
	}
}
