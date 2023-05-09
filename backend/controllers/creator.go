package controllers

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"net/http"
	"strconv"

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

func CreatorAddProduct(ctx *gin.Context) {
	email, _ := ctx.Params.Get("email")
	id, _ := ctx.Params.Get("id")
	idCreator, _ := strconv.Atoi(id)

	var newProduct dtos.ProductDTO
	ctx.ShouldBindJSON(&newProduct)

	service := services.NewProductService()
	result := service.CreateProduct(&newProduct, email, idCreator)

	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": result})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}

func CreatorAddAfiliate(ctx *gin.Context) {
	email, _ := ctx.Params.Get("email")
	id, _ := ctx.Params.Get("id")
	idCreator, _ := strconv.Atoi(id)

	var newProduct dtos.AfiliatedDTO
	ctx.ShouldBindJSON(&newProduct)

	service := services.NewAfiliatedService()
	result := service.AddAfiliate(&newProduct, email, idCreator)

	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": result})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}
