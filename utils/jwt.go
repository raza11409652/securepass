package utils

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("JWT_SECRET_TOKEN_KEY")

func GenerateSessionToken(email string, id string) string {
	var expirationTime = time.Now().Add(24 * time.Hour)
	// fmt.Print(expirationTime)
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
func TokenValid(c *gin.Context) (error, Claims) {
	tokenString := ExtractToken(c)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return err, *claims
	}
	if token.Valid {

	}
	return nil, *claims
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == "" {
		return ""
	}
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
