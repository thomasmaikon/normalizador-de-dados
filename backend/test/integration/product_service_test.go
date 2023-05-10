package integration

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()
	productService := services.NewProductService()

	login := dtos.LoginDTO{
		Email:    "EmailProductTest",
		Password: "EmailProductTest",
	}

	user := dtos.UserDTO{
		Name:  "NameProductTest",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "CreatorProductTest",
	}

	product := dtos.ProductDTO{
		Description: "Simple Product test",
		Price:       "8.5",
	}

	userOtput, _ := userService.CreateUser(user)
	creatorService.CreateNewCreator(&creator, userOtput.UserId)
	creatorService.GetCreator(userOtput.UserId)
	output := productService.CreateProduct(&product, userOtput.UserId)
	if output != nil {
		t.Fatal("Not expected error")
	}
}

func TestGetAllProducts(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()
	productService := services.NewProductService()

	login := dtos.LoginDTO{
		Email:    "EmailProductTest2",
		Password: "EmailProductTest",
	}

	user := dtos.UserDTO{
		Name:  "NameProductTest2",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "CreatorProductTest2",
	}

	products := []dtos.ProductDTO{
		{
			Description: "Simple Product test2",
			Price:       "8.5",
		},
		{
			Description: "Simple Product test3",
			Price:       "9.5",
		},
		{
			Description: "Simple Product test4",
			Price:       "10.5",
		},
	}

	userOtput, _ := userService.CreateUser(user)
	creatorService.CreateNewCreator(&creator, userOtput.UserId)
	creatorOutput, _ := creatorService.GetCreator(userOtput.UserId)

	for _, product := range products {
		output := productService.CreateProduct(&product, userOtput.UserId)
		if output != nil {
			t.Fatal("Not expected error when add product")
		}
	}

	productOuput, _ := productService.GetAllProducts(creatorOutput.CreatorId)
	if len(productOuput) != len(products) {
		t.Fatal("Not expected error")
	}
}
