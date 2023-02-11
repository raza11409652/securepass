package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/controllers"
)

func ContentRoutes(router *gin.RouterGroup) {
	router.POST("/", controllers.NewContentPost)
}
