package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/routers"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	routers.AppRouter(v1.Group("/api"))
	r.Run()
}
