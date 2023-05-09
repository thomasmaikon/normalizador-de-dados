package controllers

import (
	"hubla/desafiofullstack/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NormalizeData(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		// Lida com o erro
		ctx.String(http.StatusBadRequest, "Faild to getting file")
		return
	}

	id := ctx.GetString("userID")
	userId, _ := strconv.Atoi(id)

	service := services.NewHistoryService()
	result := service.AddHistoricalTransactions(file, userId)

	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": result})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Info:": "Data updated"})
	}
}
