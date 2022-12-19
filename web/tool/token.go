package tool

import (
	"github.com/deatil/go-cryptobin/cryptobin/crypto"
	"time"
)

type Token struct {
	Exp   int64  `json:"exp"`
	Token string `json:"encoded"`
}

//密钥
var secret = "read@book@isGood"

// TokenExpireDuration token过期时间
const TokenExpireDuration = time.Hour * 24

func GenToken(userName string) Token {
	return Token{
		Exp:   time.Now().Add(TokenExpireDuration).Unix(),
		Token: createToken(userName),
	}
}

func UnToken(encoded string) string {
	return crypto.
		FromBase64String(encoded).
		SetKey(secret).
		Aes().
		ECB().
		PKCS7Padding().
		Decrypt().
		ToString()
}

func createToken(userName string) string {
	// 加密
	return crypto.
		FromString(userName).
		SetKey(secret).
		Aes().
		ECB().
		PKCS7Padding().
		Encrypt().
		ToBase64String()
}
