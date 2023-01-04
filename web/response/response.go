package response

import "github.com/gin-gonic/gin"

func Success(message string, data any) gin.H {
	return gin.H{
		"code":    200,
		"message": message,
		"data":    data,
	}
}

func Fail(code int, message string) gin.H {
	return gin.H{
		"code":    code,
		"message": message,
	}
}
