package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/controllers"
)

func AuthRoutes(router *gin.RouterGroup) {

	// router.POST("/login", controllers.LoginMethod)
	router.POST("/register", controllers.ProfileRegistration)
}
