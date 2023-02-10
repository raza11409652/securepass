package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/controllers"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetUsers)
}
