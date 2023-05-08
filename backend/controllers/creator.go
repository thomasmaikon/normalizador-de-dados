package controllers

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewCreator(ctx *gin.Context) {
	email, _ := ctx.Params.Get("email")

	var newCreator dtos.CreatorDTO
	ctx.BindJSON(&newCreator)

	service := services.NewCreatorSerivce()
	result := service.CreateNewCreator(&newCreator, email)
	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": result})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{}) 
	}
}
