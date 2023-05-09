package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type IProductService interface {
	CreateProduct(newProduct *dtos.ProductDTO, email string, idCreator int) *dtos.ValidationDTO
	FindProduct(description string, creatorId int) (*entitys.Product, error)
}

type productService struct {
	productRepository repositorys.IProductRepository
}

func NewProductService() IProductService {
	return &productService{
		productRepository: repositorys.NewProductRepository(),
	}
}

func (service *productService) CreateProduct(newProduct *dtos.ProductDTO, email string, idCreator int) *dtos.ValidationDTO {
	isCreated, err := service.productRepository.CreateNewProduct(newProduct, email, idCreator)
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

func (service *productService) FindProduct(description string, creatorId int) (*entitys.Product, error) {
	return service.productRepository.Find(description, creatorId)
}