package router

import (
	"github.com/gin-gonic/gin"
	login "go_template/web/login/service"
)

func router(route *gin.Engine) {

	config := route.Group("/template/v1/")
	{

		// 添加权限
		config.GET("login", login.Login)
	}

}
