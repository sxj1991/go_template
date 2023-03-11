package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_template/share"
	"go_template/web/response"
	"go_template/web/tool"
	"net/http"
	"time"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//获取请求头token
		var authHeader = c.Request.Header.Get(share.TokenName)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Fail(401, "token认证失败!"))
			c.Abort()
			return
		}

		//获取用户信息(获取缓存token没有则重新登录)
		cacheByte := tool.GetCacheStruct(authHeader)

		if cacheByte != nil {
			token := checkToken(cacheByte)
			if token == nil {
				c.JSON(http.StatusUnauthorized, response.Fail(401, "token过期!"))
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, response.Fail(401, "token认证失败!"))
			c.Abort()
			return
		}

		// 解析JWT的函数来解析它
		username, _ := tool.UnToken(authHeader)
		if username == "" {
			c.JSON(http.StatusUnauthorized, response.Fail(401, "token认证失败!"))
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(authHeader, username)
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Next()
	}
}

func checkToken(cacheByte []byte) *tool.Token {
	token := tool.Token{}
	json.Unmarshal(cacheByte, &token)
	expire := time.Unix(token.Exp, 0)
	if expire.Before(time.Now()) {
		tool.RemoveCache(token.Token)
		return nil
	} else {
		expire.Add(share.TokenExpireDuration)
		token.Exp = expire.Unix()
		tool.SetCacheKey(token.Token, token)
		return &token
	}

}
