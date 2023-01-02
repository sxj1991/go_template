package tool

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"go_template/setting"
	"go_template/share"
	"io"
	"time"
)

type Token struct {
	Exp   int64  `json:"exp"`
	Token string `json:"encoded"`
	Id    string `json:"id"`
}

//密钥
var secret = []byte("read@book@isGood")

func GenToken(userName string) Token {
	token, _ := createToken(userName)

	uuid, err := uuid.NewUUID()
	if err != nil {
		setting.Log.Error(err)
	}

	return Token{
		Exp:   time.Now().Add(share.TokenExpireDuration).Unix(),
		Token: token,
		Id:    uuid.String(),
	}
}

func UnToken(encoded string) (string, error) {
	return string(aesDecryptCFB([]byte(encoded), secret)), nil
}

func createToken(userName string) (string, error) {
	return base64.StdEncoding.EncodeToString(aesEncryptCFB([]byte(userName), secret)), nil
}

// =================== CFB ======================
func aesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func aesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
