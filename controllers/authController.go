package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/models"
	"github.com/raza11409652/securepass/service"
)

func ProfileRegistration(c *gin.Context) {
	var u = models.UserModel{}
	result := service.InsertNew(u)
	c.JSONP(200, gin.H{"user": result})
}

func LoginMethod(c *gin.Context) {
	c.JSONP(200, gin.H{"error": "false"})
}
