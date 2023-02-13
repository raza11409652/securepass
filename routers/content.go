package routers

import (
	"github.com/gorilla/mux"
	"github.com/raza11409652/securepass/controllers"
)

// func ContentRoutes(router *gin.RouterGroup) {

// 	router.GET("/:id", controllers.ViewContentPost)
// 	router.Use(middleware.JwtAuthMiddleware())
// 	router.POST("/", controllers.NewContentPost)
// }

func RouterContent(r *mux.Router) {
	routes := r.PathPrefix("/content").Subrouter()
	routes.HandleFunc("/", controllers.NewContentPost)
	routes.HandleFunc("/:id", controllers.RegisterAccount).Methods("GET")
}
