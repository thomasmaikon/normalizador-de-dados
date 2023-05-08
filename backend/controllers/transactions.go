package controllers

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NormalizeData(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		// Lida com o erro
		ctx.String(http.StatusBadRequest, "Faild to getting file")
		return
	}

	log.Println(file.Filename)

	src, err := file.Open()

	scanner := bufio.NewScanner(src)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	ctx.JSON(http.StatusOK, gin.H{"Info:": "Data updated"})
}
