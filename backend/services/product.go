package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/exceptions"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type IProductService interface {
	CreateProduct(newProduct *dtos.ProductDTO, userId int) *dtos.ValidationDTO
	FindProduct(description string, creatorId int) (*entitys.Product, *dtos.ValidationDTO)
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
	if err != nil || !isCreated {
		log.Println(err)
		return &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeCreateProduct,
			Message: exceptions.ErrorMessageCreateProduct,
		}
	}

	return nil
}

func (service *productService) FindProduct(description string, creatorId int) (*entitys.Product, *dtos.ValidationDTO) {
	product, err := service.productRepository.Find(description, creatorId)
	if err != nil || product.ID == 0 {
		return nil, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeFindProduct,
			Message: exceptions.ErrorMessageFindProduct,
		}
	}

	return product, nil
}

func (service *productService) GetAllProducts(userId int) ([]*models.ProductModel, *dtos.ValidationDTO) {
	creator, validationDTO := service.creatorService.GetCreator(userId)
	if validationDTO != nil {
		return nil, validationDTO
	}

	products, err := service.productRepository.GetAll(creator.CreatorId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeFaildNotFoundALlProducts,
			Message: exceptions.ErrorMessageFaildNotFoundALlProducts,
		}
	}

	return products, nil
}
