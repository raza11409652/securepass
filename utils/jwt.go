package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("JWT_SECRET_TOKEN_KEY")
var expirationTime = time.Now().Add(24 * time.Hour)

func GenerateSessionToken(email string, id string) string {
	fmt.Print(expirationTime)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email: email,
		Id:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}
