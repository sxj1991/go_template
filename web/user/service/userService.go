package user

import (
	"github.com/gin-gonic/gin"
	"go_template/ssh"
	"go_template/web/response"
	"net/http"
)

type user map[string]any

func Users(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("响应成功", user{"name": "zhangsan", "age": 12, "verify": true}))
}

func SSH(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("响应成功",
		map[string]interface{}{"ssh返回指令结果集:": ssh.RemoteSSH()}))
}
