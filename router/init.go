package router

import (
	"github.com/gin-gonic/gin"
	"go_template/setting"
	"go_template/web/tool"
)

type ConfigRoute struct {
}

func (c *ConfigRoute) WebRouter() *gin.Engine {
	config := new(setting.Config)

	var route = gin.New()

	route.Use(gin.Recovery())

	route.Use(config.LoggerToFile())

	//根据配置进行设置跨域
	if setting.CONF.HttpServer.AllowCrossDomain {
		route.Use(tool.Next())
	}
	//初始化路由
	router(route)

	return route
}
