package controllers

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"hubla/desafiofullstack/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateLogin(ctx *gin.Context) {

	var loginDTO dtos.LoginDTO

	ctx.BindJSON(&loginDTO)

	service := services.NewUserService()

	result, validateDTO := service.GetUser(&loginDTO)
	if validateDTO != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"Info:": validateDTO})
	} else {
		token, err := utils.NewAuth().GenerateTokenJWT(result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			auth := "Bearer " + token
			ctx.Writer.Header().Set("Authorization", auth)
			ctx.JSON(http.StatusAccepted, gin.H{"token": token})
		}
	}
}
