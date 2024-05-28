package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type EncrypterService struct {
	MySecret string
}

func NewEncrypterService(mysecret string) *EncrypterService {
	return &EncrypterService{
		MySecret: mysecret,
	}
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func (es EncrypterService) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(es.MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func (es EncrypterService) Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(es.MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
