package controllers

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
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
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}
