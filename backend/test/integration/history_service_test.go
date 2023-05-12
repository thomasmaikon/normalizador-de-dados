package integration

import (
	"fmt"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
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

func TestGetAllHistorical(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()
	productService := services.NewProductService()
	afiliateService := services.NewAfiliatedService()
	historicalService := services.NewHistoricalService()

	login := dtos.LoginDTO{
		Email:    "HistoricalLogin3",
		Password: "HistoricalLogin3",
	}

	user := dtos.UserDTO{
		Name:  "UserHistorical3",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "CreatorHistorical",
	}

	products := []dtos.ProductDTO{
		{
			Description: "Simple 7",
			Price:       "1",
		},
		{
			Description: "Simple 8",
			Price:       "2",
		},
		{
			Description: "Simple 9",
			Price:       "3",
		},
	}

	afiliates := []dtos.AfiliatedDTO{
		{
			Name: "People7",
		},
		{
			Name: "People8",
		},
		{
			Name: "People9",
		},
	}

	transactionsDescriptions := map[int]string{
		1: "Venda produtor",
		2: "Venda afiliado",
		3: "Comissao paga",
		4: "Comissao recebida",
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
			IdProduct:         productsOutput[1].Id,
			Value:             0000050000,
			IdAfiliated:       afiliatesOutput[1].Id,
		},
		{
			IdTransactionType: 2,
			IdCreator:         creatorOutput.CreatorId,
			IdProduct:         productsOutput[2].Id,
			Value:             0000155000,
			IdAfiliated:       afiliatesOutput[2].Id,
		},
	}

	for _, historical := range historcals {
		validationDTO := historicalService.Add(&historical)
		if validationDTO != nil {
			t.Fatal("Err when add historical row, it is not expected")
		}
	}

	outputExpected := []*models.HistoricalModelWithOutJoins{
		{
			AfiliateName:           afiliatesOutput[0].Name,
			AfiliateId:             afiliatesOutput[0].Id,
			ProductDescription:     products[0].Description,
			TransactionDescription: transactionsDescriptions[historcals[0].IdTransactionType],
			Value:                  historcals[0].Value,
		},
		{
			AfiliateName:           afiliatesOutput[1].Name,
			AfiliateId:             afiliatesOutput[1].Id,
			ProductDescription:     products[1].Description,
			TransactionDescription: transactionsDescriptions[historcals[1].IdTransactionType],
			Value:                  historcals[1].Value,
		},
		{
			AfiliateName:           afiliatesOutput[2].Name,
			AfiliateId:             afiliatesOutput[2].Id,
			ProductDescription:     products[2].Description,
			TransactionDescription: transactionsDescriptions[historcals[2].IdTransactionType],
			Value:                  historcals[2].Value,
		},
	}
	historicalsOutput, _ := historicalService.GetAllHistorical(userOutput.UserId)

	if len(outputExpected) != len(historicalsOutput) {
		t.Fatal("Err when compare historical rows, expected it to be the same")
	}

	for index, output := range historicalsOutput {
		if *output != *outputExpected[index] {
			fmt.Println(*output)
			fmt.Println(*outputExpected[index])
			t.Fatalf("Err when compare historical rows, expected it to be the same")
		}
	}

}

func TestGetAllHistoricalFromAfiliate(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()
	productService := services.NewProductService()
	afiliateService := services.NewAfiliatedService()
	historicalService := services.NewHistoricalService()

	login := dtos.LoginDTO{
		Email:    "HistoricalLogin4",
		Password: "HistoricalLogin4",
	}

	user := dtos.UserDTO{
		Name:  "UserHistorical4",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "CreatorHistorical4",
	}

	products := []dtos.ProductDTO{
		{
			Description: "Simple 10",
			Price:       "1",
		},
		{
			Description: "Simple 11",
			Price:       "2",
		},
		{
			Description: "Simple 12",
			Price:       "3",
		},
	}

	afiliates := []dtos.AfiliatedDTO{
		{
			Name: "People10",
		},
	}

	/* transactionsDescriptions := map[int]string{
		1: "Venda produtor",
		2: "Venda afiliado",
		3: "Comissao paga",
		4: "Comissao recebida",
	} */

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
			IdProduct:         productsOutput[1].Id,
			Value:             0000050000,
			IdAfiliated:       afiliatesOutput[0].Id,
		},
		{
			IdTransactionType: 2,
			IdCreator:         creatorOutput.CreatorId,
			IdProduct:         productsOutput[2].Id,
			Value:             0000155000,
			IdAfiliated:       afiliatesOutput[0].Id,
		},
	}

	historicalAfiliateExpected := &dtos.HistoricalCompleteAfiliateDTO{
		AfiliateHistoricals: &[]models.HistoricalModelWithOutJoins{
			models.HistoricalModelWithOutJoins{
				AfiliateName:           "People10",
				AfiliateId:             0,
				ProductDescription:     "Simple 7",
				TransactionDescription: "Venda produtor",
				Value:                  0000012750,
			},
			models.HistoricalModelWithOutJoins{
				AfiliateName:           "People10",
				AfiliateId:             0,
				ProductDescription:     "Simple 8",
				TransactionDescription: "Comissao recebida",
				Value:                  0000050000,
			},
			models.HistoricalModelWithOutJoins{
				AfiliateName:           "People10",
				AfiliateId:             0,
				ProductDescription:     "Simple 9",
				TransactionDescription: "Venda afiliado",
				Value:                  0000155000,
			},
		},
		Amount: 0000050000,
	}

	for _, historical := range historcals {
		validationDTO := historicalService.Add(&historical)
		if validationDTO != nil {
			t.Fatal("Err when add historical row, it is not expected")
		}
	}

	historicalsOutput, _ := historicalService.GetAfiliateHistorical(userOutput.UserId, afiliatesOutput[0].Id)

	if len(*historicalsOutput.AfiliateHistoricals) != len(*historicalAfiliateExpected.AfiliateHistoricals) ||
		historicalsOutput.Amount != historicalAfiliateExpected.Amount {
		t.Fatal("Unexpected error, objects were supposed to be the same")
	}
}
