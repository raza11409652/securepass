package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/common"
	"github.com/raza11409652/securepass/models"
	"github.com/raza11409652/securepass/service"
	"github.com/raza11409652/securepass/utils"
)

func NewVaultCreation(ctx *gin.Context) {
	var post models.SecureVault
	_, claims := utils.TokenValid(ctx)
	if claims.Role == "ADMIN" {

	} else {
		post.MasterPassword = true
	}

	if err := common.Bind(ctx, &post); err != nil {
		ctx.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErrors := validate.Struct(&post)
	if validationErrors != nil {
		commonError := common.NewError("error_message", validationErrors)
		ctx.JSONP(http.StatusBadRequest, commonError)
		return
	}
	post.User = claims.Id
	hashCode := utils.RandomStr(32)
	post.VaultKey = hashCode
	post.Notes = utils.EncryptAES(hashCode, post.Notes)
	if post.Type == "PASSWORD" {
		if post.Url == "" || !utils.IsValidUrl(post.Url) {
			ctx.AbortWithStatusJSON(400, gin.H{"message": "URL is required"})
			return
		}
		if post.Username == "" {
			ctx.AbortWithStatusJSON(400, gin.H{"message": "Username is required"})
			return
		}
		if post.Password == "" {
			ctx.AbortWithStatusJSON(400, gin.H{"message": "Password is required"})
			return
		}
		post.Username = utils.EncryptAES(hashCode, post.Username)
		post.Password = utils.EncryptAES(hashCode, post.Password)
	}
	result := service.InsertNewVault(post)
	ctx.JSON(200, result)
	return
}
