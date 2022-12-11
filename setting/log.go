package setting

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

// 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var Log *logrus.Logger

func logger() {
	//获取文件路径
	var filenameOnly = strings.TrimSuffix(CONF.LogFile, path.Base(CONF.LogFile))
	exists, _ := pathExists(filenameOnly)
	//如何日志文件夹没有创建则创建日志目录
	if !exists {
		_ = os.MkdirAll(filenameOnly, os.ModePerm)
	}
	//写入文件
	_, err := os.OpenFile(CONF.LogFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		logrus.Fatalln(err)
	}

	//实例化
	Log = logrus.New()
	//设置日志级别
	Log.SetLevel(logrus.InfoLevel)

	//设置日志格式
	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//日志分割
	path := CONF.LogFile
	logConfig, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),                            //软连接
		rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),      //保持7天的日志
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), //每一天生成一个文件
		//rotatelogs.WithRotationCount(7),                          //保持7个文件与WithMaxAge只能选其一
	)
	//按日志级别分割
	pathDebug := CONF.LogErrFile
	logWarn, _ := rotatelogs.New(
		pathDebug+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(pathDebug), //软连接
		//rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),      //保持7天的日志
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), //每一天生成一个文件
		rotatelogs.WithRotationCount(7),                          //保持7个文件与WithMaxAge只能选其一
	)
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		// 为不同级别设置不同的输出目的
		logrus.WarnLevel:  logWarn,
		logrus.ErrorLevel: logWarn,
	}, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//添加钩子
	Log.AddHook(lfHook)

	gin.DefaultWriter = io.MultiWriter(logConfig)

	if CONF.Debug {
		//打印到控制台
		Log.SetOutput(io.MultiWriter(os.Stdout, logConfig))
	} else {
		Log.SetOutput(logConfig)
	}

}

func (c *Config) LoggerToFile() gin.HandlerFunc {
	logger := Log
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
