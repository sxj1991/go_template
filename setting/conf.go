package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go_template/share"
	"go_template/web/tool"
	"os"
	"path"
	"regexp"
	"strings"
)

type Config struct {
	//启动模式，是否debug
	Debug bool `yaml:"Debug"`
	//web服务配置
	HttpServer httpServer `yaml:"HttpServer"`
	//日志文件目录
	LogFile string `yaml:"LogFile"`
	//日志错误信息文件目录
	LogErrFile string `yaml:"LogErrFile"`
	//admin
	User struct {
		Name       string `yaml:"Name"`
		Password   string `yaml:"Password"`
		VerifyCode string `yaml:"VerifyCode"`
	} `yaml:"User"`
}

type httpServer struct {
	//项目启动的端口 前端端口
	Port string `yaml:"Port"`
	//开启跨域
	AllowCrossDomain bool `yaml:"AllowCrossDomain"`
}

var CONF *Config

/**
  读取配置文件，初始化服务配置
  初始化配置和日志
*/
func init() {
	conf()
	logger()
}

func conf() {
	//获取配置文件路径
	fileName := os.Getenv(share.CONFIGURE)
	if "" == fileName {
		logrus.Panicf(share.ErrorsConfigNotExists)
	}
	filenameWithSuffix := path.Base(fileName)
	v := viper.New()
	v.SetConfigName(strings.TrimSuffix(filenameWithSuffix, path.Ext(filenameWithSuffix))) // 设置文件名称（无后缀）
	v.AddConfigPath(strings.TrimSuffix(fileName, filenameWithSuffix))                     // 设置文件所在路径

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("请检查配置文件路径")
		} else {
			panic("配置文件出错")
		}
	}
	var conf Config

	//监控配置和重新获取配置
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		_ = v.Unmarshal(&conf)
		hasDecrypt(&conf)
		CONF = &conf
		logrus.Warnf("配置信息已修改：%v\n", conf)

	})
	_ = v.Unmarshal(&conf)
	hasDecrypt(&conf)
	CONF = &conf

}

/** 判断是否存在加密 并解密内容*/
func hasDecrypt(conf *Config) {
	if strings.Contains(conf.User.Name, "ENC") {
		conf.User.Name = configDecrypt(conf.User.Name)
	}

	if strings.Contains(conf.User.Password, "ENC") {
		conf.User.Password = configDecrypt(conf.User.Password)
	}

	if strings.Contains(conf.User.VerifyCode, "ENC") {
		conf.User.VerifyCode = configDecrypt(conf.User.VerifyCode)
	}
}

func configDecrypt(encrypt string) (decrypt string) {
	r, _ := regexp.Compile("\\(.+?\\)")
	aes, err := tool.DecryptByAes(r.FindString(encrypt)[1 : len(r.FindString(encrypt))-1])
	if nil != err {
		return ""
	}
	return string(aes)
}
