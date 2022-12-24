package login

import (
	"github.com/gin-gonic/gin"
	"go_template/web/tool"
	"net/http"
)

type login struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
}

func Login(c *gin.Context) {
	var l login
	if err := c.ShouldBind(&l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 401,
			"msg":  err.Error(),
		})
		return
	}

	if err := l.checkUser(); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  err.Error(),
		})
		return
	}

	token := tool.GenToken(l.UserName)
	tool.SetCacheKey(token.Token, token)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  token.Token,
	})
}

func (l *login) checkUser() error {
	//TODO 暂时没有用户数据库表 只不过一个基本校验
	if l.UserName == "zhangsan" && l.Password == "123456" && l.VerifyCode == "qaz" {
		return nil
	}
	return l
}

func (l *login) Error() string {
	return "登录验证错误"
}
