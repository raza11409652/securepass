package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/controllers"
	"github.com/raza11409652/securepass/middleware"
)

func VaultRoutes(router *gin.RouterGroup) {
	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/", controllers.NewVaultCreation)
}
