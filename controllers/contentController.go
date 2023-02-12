package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/common"
	"github.com/raza11409652/securepass/models"
	"github.com/raza11409652/securepass/service"
	"github.com/raza11409652/securepass/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	if post.MaxViewAllowed < 1 {
		post.MaxViewAllowed = 1
	}
	// resultX := utils.DecryptAES(hashCode, encryption)
	result := service.InsertNewContent(post)
	ctx.JSON(200, gin.H{"id": result.InsertedID, "url": URL})
	return
}

func ViewContentPost(ctx *gin.Context) {
	var content models.SecureDataModel
	id := ctx.Param("id")
	key := ctx.Query("key")
	if len(key) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Key is required"})
		return
	}
	ip := ctx.ClientIP()
	device := ctx.Request.UserAgent()
	referrer := ctx.Request.Referer()
	fmt.Print(ip, device, referrer)
	err := service.GetContentByFilter(bson.M{"finderKey": id}).Decode(&content)
	// ctx.JSONP(200, data)
	if err == mongo.ErrNoDocuments {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	flag := utils.CompareBcryptHash(key, content.HashCode)
	if !flag {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Invalid key passed"})
		return
	}
	hCount := service.GetContentHistoryCount(bson.M{"content": content.ID})
	if hCount >= content.MaxViewAllowed {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "URL expired"})
		return
	}
	// Insert history
	var history models.SecureDataHistoryModel
	history.CreatedAt, _ = utils.GetCurrentTime()
	history.UpdatedAt, _ = utils.GetCurrentTime()
	history.IpAddress = ip
	history.Device = device
	history.Referrer = referrer
	history.Content = content.ID
	result := service.InsertNewContentHistory(history)
	content.Content = utils.DecryptAES(key, content.Content)
	ctx.JSON(http.StatusOK, gin.H{"h": result, "data": content})
	return
}
