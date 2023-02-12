package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err, claims := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "UnAuthorized"})
			c.Abort()
			return
		}
		fmt.Print(claims)
		// c.SetCookie()
		c.Next()
	}
}
