package router

import (
	"github.com/gin-gonic/gin"
	userService "go_template/web/user"
)

func router(route *gin.Engine) {

	config := route.Group("/template/v1/")
	{
		// 用户登录
		config.GET("login", userService.UserLogin)

	}

}
