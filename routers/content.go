package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/controllers"
	"github.com/raza11409652/securepass/middleware"
)

func ContentRoutes(router *gin.RouterGroup) {
	router.GET("/:id", controllers.ViewContentPost)
	router.Use(middleware.JwtAuthMiddleware())
	router.POST("/", controllers.NewContentPost)
}
