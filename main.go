package main

import (
	"go_template/router"
	"go_template/setting"
)

func main() {
	//初始化
	route := new(router.ConfigRoute)
	webRouter := route.WebRouter()
	//启动服务
	setting.Log.Infof("go template start port:%v", setting.CONF.HttpServer.Port)
	_ = webRouter.Run(setting.CONF.HttpServer.Port)
}
