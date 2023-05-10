package integration

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"testing"
)

func TestAddTransactionsInHistorical(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()
	productService := services.NewProductService()
	afiliateService := services.NewAfiliatedService()
	historicalService := services.NewHistoricalService()

	login := dtos.LoginDTO{
		Email:    "HistoricalLogin",
		Password: "HistoricalLogin",
	}

	user := dtos.UserDTO{
		Name:  "UserHistorical",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "CreatorHistorical",
	}

	products := []dtos.ProductDTO{
		{
			Description: "Simple 1",
			Price:       "5",
		},
		{
			Description: "Simple 2",
			Price:       "5",
		},
		{
			Description: "Simple 3",
			Price:       "5",
		},
	}

	afiliates := []dtos.AfiliatedDTO{
		{
			Name: "People1",
		},
		{
			Name: "People2",
		},
		{
			Name: "People3",
		},
	}

	userOutput, _ := userService.CreateUser(user)
	creatorService.CreateNewCreator(&creator, userOutput.UserId)
	creatorOutput, _ := creatorService.GetCreator(userOutput.UserId)

	for _, product := range products {
		productService.CreateProduct(&product, userOutput.UserId)
	}

	for _, afiliate := range afiliates {
		afiliateService.AddAfiliate(&afiliate, userOutput.UserId)
	}

	productsOutput, _ := productService.GetAllProducts(userOutput.UserId)
	afiliatesOutput, _ := afiliateService.GetAllAfiliates(userOutput.UserId)

	historcals := []dtos.HistoryCompleteDTO{
		{
			IdTransactionType: 1,
			IdCreator:         creatorOutput.CreatorId,
			IdProduct:         productsOutput[0].Id,
			Value:             0000012750,
			IdAfiliated:       afiliatesOutput[0].Id,
		},
		{
			IdTransactionType: 4,
			IdCreator:         creatorOutput.CreatorId,
			IdProduct:         productsOutput[2].Id,
			Value:             0000050000,
			IdAfiliated:       afiliatesOutput[2].Id,
		},
		{
			IdTransactionType: 2,
			IdCreator:         creatorOutput.CreatorId,
			IdProduct:         productsOutput[1].Id,
			Value:             0000155000,
			IdAfiliated:       afiliatesOutput[1].Id,
		},
	}

	for _, historical := range historcals {
		validationDTO := historicalService.Add(&historical)
		if validationDTO != nil {
			t.Fatal("Err when add historical row, it is not expected")
		}
	}
}

func TestAddTransactionsInHistoricalWithInvalidProduct(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()
	afiliateService := services.NewAfiliatedService()
	historicalService := services.NewHistoricalService()

	login := dtos.LoginDTO{
		Email:    "HistoricalLogin2",
		Password: "HistoricalLogin2",
	}

	user := dtos.UserDTO{
		Name:  "UserHistorical2",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "CreatorHistorical2",
	}

	afiliates := []dtos.AfiliatedDTO{
		{
			Name: "People4",
		},
		{
			Name: "People5",
		},
		{
			Name: "People6",
		},
	}

	userOutput, _ := userService.CreateUser(user)
	creatorService.CreateNewCreator(&creator, userOutput.UserId)
	creatorOutput, _ := creatorService.GetCreator(userOutput.UserId)

	for _, afiliate := range afiliates {
		afiliateService.AddAfiliate(&afiliate, userOutput.UserId)
	}

	afiliatesOutput, _ := afiliateService.GetAllAfiliates(userOutput.UserId)

	historcalRow := dtos.HistoryCompleteDTO{
		IdTransactionType: 2,
		IdCreator:         creatorOutput.CreatorId,
		IdProduct:         59555,
		Value:             0000155000,
		IdAfiliated:       afiliatesOutput[1].Id,
	}

	validationDTO := historicalService.Add(&historcalRow)
	if validationDTO == nil {
		t.Fatal("Expected error, because the product_id 59555 not be inserted")
	}
}
