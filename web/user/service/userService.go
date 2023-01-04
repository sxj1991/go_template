package user

import (
	"github.com/gin-gonic/gin"
	"go_template/web/response"
	"net/http"
)

type user map[string]any

func Users(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("响应成功", user{"name": "zhangsan", "age": 12, "verify": true}))
}
