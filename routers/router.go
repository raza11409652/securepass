package routers

import "github.com/gin-gonic/gin"

func AppRouter(router *gin.RouterGroup) {
	AuthRoutes(router.Group("/auth"))
	UserRoutes(router.Group("/users"))
	ContentRoutes(router.Group("/secure-content"))

}
