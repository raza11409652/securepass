package utils

import (
	"math/rand"

	"github.com/google/uuid"
)

func GenerateUniqueId() string {
	return uuid.New().String()
}

// define the given charset, char only

// n is the length of random string we want to generate
func RandomStr(n int) string {
	var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// func GenerateSonyFlakeId() string {
// 	result, error := sonyflake.NewSonyflake().NextID()
// 	if error != nil {
// 		log.Fatal(error)
// 	}
// 	return string(result)
// }
