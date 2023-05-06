package controllers

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateLogin(ctx *gin.Context) {

	var loginDTO dtos.LoginDTO

	ctx.BindJSON(&loginDTO)

	service := services.NewLoginService()

	result := service.ValidateCredential(loginDTO)
	if result != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"Info:": result})
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"token": "test TOken"})
	}
}
