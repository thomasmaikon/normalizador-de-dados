package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type IProductService interface {
	CreateProduct(newProduct *dtos.ProductDTO, email string, idCreator int) *dtos.ValidationDTO
}

type productService struct {
	repositorys.IProductRepository
}

func NewProductService() IProductService {
	return &productService{
		IProductRepository: repositorys.NewProductRepository(),
	}
}

func (service *productService) CreateProduct(newProduct *dtos.ProductDTO, email string, idCreator int) *dtos.ValidationDTO {
	isCreated, err := service.IProductRepository.CreateNewProduct(newProduct, email, idCreator)
	if err != nil {
		log.Println(err)
		return &dtos.ValidationDTO{
			Code:    9,
			Message: "Error when add new product",
		}
	} else if !isCreated {
		return &dtos.ValidationDTO{
			Code:    9,
			Message: "Invalid creator data to add product",
		}
	}

	return nil
}
