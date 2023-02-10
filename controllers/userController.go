package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/service"
)

func GetUsers(c *gin.Context) {
	records := service.GetUsers()
	c.JSON(200, gin.H{"records": records})
	return
}
