package router

import (
	"github.com/gin-gonic/gin"
	"go_template/setting"
	login "go_template/web/login/service"
	"go_template/web/middleware"
	user "go_template/web/user/service"
)

func router(route *gin.Engine) {
	// 基础路径请求
	config := route.Group("/template/v1/")
	{

		// 登录
		config.POST("login", login.Login)

		// ssh 连接服务器
		config.GET("ssh", user.SSH)

		// 认证中间件
		config.Use(middleware.AuthMiddleware())

		// 获取用户
		config.GET("user", user.Users)
	}

	// 是否开启性能分析工具
	if setting.CONF.Debug {
		profiling(route)
	}

}
