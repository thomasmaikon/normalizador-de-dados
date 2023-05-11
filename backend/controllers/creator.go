package controllers

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNewCreator(ctx *gin.Context) {
	id := ctx.GetString("userID")
	userId, _ := strconv.Atoi(id)

	var newCreator dtos.CreatorDTO
	ctx.BindJSON(&newCreator)

	service := services.NewCreatorSerivce()
	result := service.CreateNewCreator(&newCreator, userId)
	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": result})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}

func GetCreator(ctx *gin.Context) {
	id := ctx.GetString("userID")
	userId, _ := strconv.Atoi(id)

	serviceCreator := services.NewCreatorSerivce()
	result, validationDTO := serviceCreator.GetCreator(userId)
	if validationDTO != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Info": validationDTO})
		return
	}

	serviceHistorical := services.NewHistoricalService()
	ammount, validationDTO := serviceHistorical.GetAmmountAtCreator(result.CreatorId)
	if validationDTO != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Info": validationDTO})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Info": result, "Amount": ammount})

}

func CreatorAddProduct(ctx *gin.Context) {
	id := ctx.GetString("userID")
	userId, _ := strconv.Atoi(id)

	var newProduct dtos.ProductDTO
	ctx.ShouldBindJSON(&newProduct)

	service := services.NewProductService()
	result := service.CreateProduct(&newProduct, userId)

	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": result})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}

func CreatorAddAfiliate(ctx *gin.Context) {
	id := ctx.GetString("userID")
	userId, _ := strconv.Atoi(id)

	var newProduct dtos.AfiliatedDTO
	ctx.ShouldBindJSON(&newProduct)

	service := services.NewAfiliatedService()
	result := service.AddAfiliate(&newProduct, userId)

	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": result})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{})
	}
}

func GetHistoricalTransactions(ctx *gin.Context) {
	id := ctx.GetString("userID")
	userId, _ := strconv.Atoi(id)

	service := services.NewHistoricalService()
	resutl, validationDTO := service.GetAllHistorical(userId)

	if validationDTO != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": validationDTO})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Info": resutl})
	}
}

func GetHistoricalTransactionsAtAfiliate(ctx *gin.Context) {
	id := ctx.GetString("userID")
	userId, _ := strconv.Atoi(id)

	paramId := ctx.Param("id")
	afiliateId, _ := strconv.Atoi(paramId)

	service := services.NewHistoricalService()
	result, validationDTO := service.GetAfiliateHistorical(userId, afiliateId)

	if validationDTO != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Info": validationDTO})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"Info": result})
	}

}
