package routers

import (
	"github.com/gorilla/mux"
	"github.com/raza11409652/securepass/controllers"
)

// func AppRouter(router *gin.RouterGroup) {
// 	// router.Use(middleware.CORS())
// 	AuthRoutes(router.Group("/auth"))
// 	UserRoutes(router.Group("/users"))
// 	ContentRoutes(router.Group("/secure-content"))
// 	TeamRoutes(router.Group("/team"))
// 	VaultRoutes(router.Group("/vault"))

// }
func RouterMUX(r *mux.Router) {
	routes := r.PathPrefix("/auth").Subrouter()
	routes.HandleFunc("/login", controllers.LoginMethod)
	routes.HandleFunc("/register", controllers.RegisterAccount)
}
