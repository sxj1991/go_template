package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg": "hello,world",
	})
}
