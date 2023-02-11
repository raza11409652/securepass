package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/raza11409652/securepass/common"
	"github.com/raza11409652/securepass/models"
	"github.com/raza11409652/securepass/service"
	"github.com/raza11409652/securepass/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

// Profile registration
func ProfileRegistration(c *gin.Context) {
	var user models.UserModel
	if err := common.Bind(c, &user); err != nil {
		// log.Panic(err)
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErrors := validate.Struct(user)
	if validationErrors != nil {
		commonError := common.NewError("error_message", validationErrors)
		c.JSONP(http.StatusBadRequest, commonError)
		return
	}
	// Check is this email already exist
	count := service.GetUserCount(bson.M{"email": user.Email})
	if count > 0 {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "Email already exist"})
		return
	}
	user.Password = utils.GenerateBcryptHash(user.Password)
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result := service.InsertNewUser(user)

	// We need to send an email which contain email verification link
	c.JSONP(200, gin.H{"user": result})
}

// Profile login
func LoginMethod(c *gin.Context) {
	var body models.LoginBody
	var user models.UserModel
	if err := common.Bind(c, &body); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErrors := validate.Struct(body)
	if validationErrors != nil {
		commonError := common.NewError("error_message", validationErrors)
		c.JSONP(http.StatusBadRequest, commonError)
		return
	}
	err := service.GetUser(bson.M{"email": body.Email}).Decode(&user)
	// fmt.Print(*userProfile)
	if err == mongo.ErrNoDocuments {
		c.JSONP(http.StatusNotFound, gin.H{"message": "Auth failed"})
		return
	}
	flag := utils.CompareBcryptHash(body.Password, user.Password)
	if !flag {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "Auth failed , Email and password mismatch"})
		return
	}
	profile := map[string]string{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"profile":   user.ProfileImage,
		"email":     user.Email,
	}
	token := utils.GenerateSessionToken(user.Email, user.ID.String())
	c.JSONP(200, gin.H{"token": token, "profile": profile})
}
