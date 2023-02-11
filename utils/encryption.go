package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func EncryptAES(key string, plaintext string) string {

	text := []byte(plaintext)
	//create a new cipher Block from the key

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plaintext))
	return base64.URLEncoding.EncodeToString(cipherText)
}

func DecryptAES(key string, ct string) string {
	cipherText, _ := base64.URLEncoding.DecodeString(ct)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	if len(cipherText) < aes.BlockSize {
		panic("Cipher text length is small")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText)
}
