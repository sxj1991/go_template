package userService

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {

	c.JSON(http.StatusOK, "登录成功")
}
