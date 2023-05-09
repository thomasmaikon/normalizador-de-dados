package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type IProductService interface {
	CreateProduct(newProduct *dtos.ProductDTO, userId int) *dtos.ValidationDTO
	FindProduct(description string, creatorId int) (*entitys.Product, error)
	GetAllProducts(userId int) ([]*models.ProductModel, *dtos.ValidationDTO)
}

type productService struct {
	productRepository repositorys.IProductRepository
	creatorService    ICreatorService
}

func NewProductService() IProductService {
	return &productService{
		productRepository: repositorys.NewProductRepository(),
		creatorService:    NewCreatorSerivce(),
	}
}

func (service *productService) CreateProduct(newProduct *dtos.ProductDTO, userId int) *dtos.ValidationDTO {

	creator, validationDTO := service.creatorService.GetCreator(userId)
	if validationDTO != nil {
		return validationDTO
	}

	isCreated, err := service.productRepository.CreateNewProduct(newProduct, userId, creator.CreatorId)
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

func (service *productService) GetAllProducts(userId int) ([]*models.ProductModel, *dtos.ValidationDTO) {
	creator, validationDTO := service.creatorService.GetCreator(userId)
	if validationDTO != nil {
		return nil, validationDTO
	}

	products, err := service.productRepository.GetAll(creator.CreatorId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    21,
			Message: "Faild when get all products",
		}
	}

	return products, nil
}
