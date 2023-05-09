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

	id, _ := ctx.Params.Get("id")
	idCreator, _ := strconv.Atoi(id)

	//log.Println(file.Filename)

	service := services.NewHistoryService()
	result := service.AddHistoryAtTransactions(file, idCreator)

	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": result})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Info:": "Data updated"})
	}
}
