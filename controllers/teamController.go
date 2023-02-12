package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raza11409652/securepass/common"
	"github.com/raza11409652/securepass/models"
	"github.com/raza11409652/securepass/service"
	"github.com/raza11409652/securepass/utils"
)

func NewTeamCreation(ctx *gin.Context) {
	_, claims := utils.TokenValid(ctx)
	user := claims.Id
	if claims.Role != "ADMIN" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Not allowed")
		return
	}
	// fmt.Print(user)
	var team models.TeamData
	if err := common.Bind(ctx, &team); err != nil {
		ctx.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validationErrors := validate.Struct(&team)
	if validationErrors != nil {
		commonError := common.NewError("error_message", validationErrors)
		ctx.JSONP(http.StatusBadRequest, commonError)
		return
	}
	if team.MasterPassword != "" {
		team.MasterPassword = utils.GenerateBcryptHash(team.MasterPassword)
	}
	team.CreatedBy = user
	team.CreatedAt, _ = utils.GetCurrentTime()
	result := service.InsertNewTeam(team)
	ctx.JSON(200, result)
	return
}

func GetTeamList(ctx *gin.Context) {
	_, claims := utils.TokenValid(ctx)
	if claims.Role != "ADMIN" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Not allowed")
		return
	}
	var teams []models.TeamData
	c := service.GetTeams(claims.Id)
	if err := c.All(context.Background(), &teams); err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, teams)
	return
}
