package tool

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go_template/share"
	"math/rand"
	"time"
)

type Token struct {
	Exp   int64  `json:"exp"`
	Token string `json:"encoded"`
}

//密钥
var secret = []byte("read@book@isGood")

func GenToken(userName string) Token {
	token, _ := createToken(userName)

	return Token{
		Exp:   time.Now().Add(share.TokenExpireDuration).Unix(),
		Token: token,
	}
}

func UnToken(encoded string) (string, error) {
	//需要用和加密时同样的方式转化成对应的字节数组
	token, err := jwt.Parse(encoded, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return "", err
	}

	if token.Claims == nil {
		return "", errors.New("token解析出错")
	}

	claims := token.Claims.(jwt.MapClaims)

	if claims["userName"] == nil {
		return "", errors.New("token解析出错")
	}

	return claims["userName"].(string), nil
}

func createToken(userName string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	claims := jwt.MapClaims{
		"userName": userName, //保存的信息
		"id":       rand.Int(),
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}
