package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/repositorys"
	"mime/multipart"
)

type IHistoricalService interface {
	Add(historyRow *dtos.HistoryCompleteDTO) *dtos.ValidationDTO
	AddHistoricalTransactions(file *multipart.FileHeader, userId int) *dtos.ValidationDTO
}

type historicalService struct {
	historyRepository     repositorys.IHistoricalRepository
	normalizedDataService InormalizeDataService
	afiliatedService      IAfiliatedService
	productService        IProductService
	creatorService        ICreatorService
}

func NewHistoricalService() IHistoricalService {
	return &historicalService{
		historyRepository:     repositorys.NewHistoricalService(),
		normalizedDataService: NewNormalizeDataService(),
		afiliatedService:      NewAfiliatedService(),
		productService:        NewProductService(),
		creatorService:        NewCreatorSerivce(),
	}
}

func (service *historicalService) Add(historyRow *dtos.HistoryCompleteDTO) *dtos.ValidationDTO {
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

func (service *historicalService) AddHistoricalTransactions(file *multipart.FileHeader, userId int) *dtos.ValidationDTO {
	normalizedData, err := service.normalizedDataService.GetNormalizedData(file)
	if err != nil {
		return &dtos.ValidationDTO{
			Code:    16,
			Message: "Faild to normalize data in file inserted",
		}
	}

	creator, validationDTO := service.creatorService.GetCreator(userId)
	if validationDTO != nil {
		return validationDTO
	}

	service.historyRepository.Begin()
	for _, historcal := range normalizedData {
		afiliate, err := service.afiliatedService.FindAfiliate(historcal.Afiliate, creator.CreatorId)
		if err != nil || afiliate.ID == 0 {
			service.historyRepository.Rollback()
			return &dtos.ValidationDTO{
				Code:    17,
				Message: "Does`not find respective afiliated",
			}
		}

		product, err := service.productService.FindProduct(historcal.ProductDescription, creator.CreatorId)
		if err != nil || product.ID == 0 {
			service.historyRepository.Rollback()
			return &dtos.ValidationDTO{
				Code:    18,
				Message: "Does`not find respective product",
			}
		}

		validationDTO := service.Add(&dtos.HistoryCompleteDTO{
			Date:              historcal.Date,
			IdCreator:         creator.CreatorId,
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
