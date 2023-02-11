package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/common"
	"github.com/raza11409652/securepass/models"
	"github.com/raza11409652/securepass/service"
	"github.com/raza11409652/securepass/utils"
)

// Save new content post and generate new URL
func NewContentPost(ctx *gin.Context) {
	var post models.SecureDataModel
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

	// 32 bit string
	hashCode := utils.RandomStr(32)
	key := strconv.Itoa((time.Now().Minute()+time.Now().Second())*10) + utils.RandomStr(10)
	// This hash code will be used for encryption
	post.HashCode = utils.GenerateBcryptHash(hashCode)
	post.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	post.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	// ctx.JSON(200, gin.H{"code": hashCode})

	//Data encryption
	encryption := utils.EncryptAES(hashCode, post.Content)
	post.Content = encryption
	post.FinderKey = key
	URL := "http://localhost:5173/secure-content/" + key + "#" + hashCode
	post.Url = URL
	// resultX := utils.DecryptAES(hashCode, encryption)
	result := service.InsertNewContent(post)
	ctx.JSON(200, gin.H{"id": result.InsertedID, "url": URL})
	return
}
