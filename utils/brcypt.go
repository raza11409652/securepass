package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Generate Bcrypt hash for password
func GenerateBcryptHash(password string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}

func CompareBcryptHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
