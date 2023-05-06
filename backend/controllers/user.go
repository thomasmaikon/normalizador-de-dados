package controllers

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"hubla/desafiofullstack/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var inputUser dtos.CreateUseDTO
	ctx.BindJSON(&inputUser)

	service := services.NewUserService()

	result := service.CreateUser(inputUser)

	if result != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": result})
	} else {
		token, err := utils.NewAuth().GenerateTokenJWT(&inputUser.Login)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			auth := "Bearer " + token
			ctx.Writer.Header().Set("Authorization", auth)
			ctx.JSON(http.StatusCreated, gin.H{"token": token})
		}
	}
}
