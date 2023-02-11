package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateBcryptHash(password string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}

func CompareBcryptHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
