package controllers

import (
	"encoding/json"
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
	team := c.Query("team")
	if team != "" {
		user.Role = "USER"
	} else {
		user.Role = "ADMIN"
	}
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

func RegisterAccount(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.UserModel
	var body models.RegisterBody
	json.NewDecoder(req.Body).Decode(&body)
	team := req.URL.Query().Get("team")
	if team != "" {
		user.Role = "USER"
	} else {
		user.Role = "ADMIN"
	}
	validationErrors := validate.Struct(body)
	if validationErrors != nil {
		commonError := common.NewError("error_message", validationErrors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commonError)
		return
	}
	count := service.GetUserCount(bson.M{"email": user.Email})
	if count > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Email already registered")
		return
	}
	user.Password = utils.GenerateBcryptHash(user.Password)
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.FirstName = body.FirstName
	user.LastName = body.LastName
	result := service.InsertNewUser(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	return

}

// Profile login
func LoginMethod(w http.ResponseWriter, req *http.Request) {
	var body models.LoginBody
	var user models.UserModel
	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(req.Body).Decode(&body)

	validationErrors := validate.Struct(body)
	if validationErrors != nil {
		commonError := common.NewError("error_message", validationErrors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(commonError)
		return
	}
	err := service.GetUser(bson.M{"email": body.Email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Auth failed")
		return
	}
	flag := utils.CompareBcryptHash(body.Password, user.Password)
	if !flag {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Password combination failed")
		return
	}
	profile := map[string]string{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"profile":   user.ProfileImage,
		"email":     user.Email,
	}
	token := utils.GenerateSessionToken(user.Email, user.ID, user.Role)
	response := map[string]any{"profile": profile, "token": token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
