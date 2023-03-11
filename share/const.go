package share

import "time"

const (
	// CONFIGURE 配置文件环境变量名
	CONFIGURE = "conf"
	// TokenName http请求header中的token的键名
	TokenName = "Authentication"
	// TokenExpireDuration token过期时间
	TokenExpireDuration = time.Minute * 30
)
