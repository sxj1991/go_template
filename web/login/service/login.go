package login

import (
	"github.com/gin-gonic/gin"
	"go_template/web/tool"
	"net/http"
)

func Login(c *gin.Context) {
	token := tool.GenToken("zhangsan")
	c.JSON(http.StatusOK, token)
}
