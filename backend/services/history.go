package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/repositorys"
	"mime/multipart"
)

type IHistoryService interface {
	Add(historyRow *dtos.HistoryCompleteDTO) *dtos.ValidationDTO
	AddHistoryAtTransactions(file *multipart.FileHeader, creatorId int) *dtos.ValidationDTO
}

type historyService struct {
	historyRepository repositorys.IHistoryRepository
	normalizedData    InormalizeDataService
	afiliatedService  IAfiliatedService
	productService    IProductService
}

func NewHistoryService() IHistoryService {
	return &historyService{
		historyRepository: repositorys.NewHistoryService(),
		normalizedData:    NewNormalizeDataService(),
		afiliatedService:  NewAfiliatedService(),
		productService:    NewProductService(),
	}
}

func (service *historyService) Add(historyRow *dtos.HistoryCompleteDTO) *dtos.ValidationDTO {
	isComplete, err := service.historyRepository.AddHistoryRow(historyRow)
	if err != nil {
		return &dtos.ValidationDTO{
			Code:    15,
			Message: "Doesn`t possible add row at file",
		}
	} else if !isComplete {
		return &dtos.ValidationDTO{
			Code:    16,
			Message: "Error when insert data in history",
		}
	}

	return nil
}

func (service *historyService) AddHistoryAtTransactions(file *multipart.FileHeader, creatorId int) *dtos.ValidationDTO {
	normalizedData, err := service.normalizedData.GetNormalizedData(file)
	if err != nil {
		return &dtos.ValidationDTO{
			Code:    16,
			Message: "Faild to normalize data in file inserted",
		}
	}

	service.historyRepository.Begin()
	for _, historcal := range normalizedData {
		afiliate, err := service.afiliatedService.FindAfiliate(historcal.Afiliate, creatorId)
		if err != nil {
			return &dtos.ValidationDTO{
				Code:    17,
				Message: "Does`not find respective afiliated",
			}
		}

		product, err := service.productService.FindProduct(historcal.ProductDescription, creatorId)
		if err != nil {
			return &dtos.ValidationDTO{
				Code:    18,
				Message: "Does`not find respective product",
			}
		}
		validationDTO := service.Add(&dtos.HistoryCompleteDTO{
			Date:              historcal.Date,
			IdCreator:         creatorId,
			Value:             historcal.Value,
			IdProduct:         product.ID,
			IdAfiliated:       afiliate.ID,
			IdTransactionType: historcal.IdTransactionType,
		})

		if validationDTO != nil {
			service.historyRepository.Rollback()
			return &dtos.ValidationDTO{
				Code:    19,
				Message: "Faild to add historicals, values are conflicted",
			}
		}
	}
	service.historyRepository.Commit()
	return nil
}
