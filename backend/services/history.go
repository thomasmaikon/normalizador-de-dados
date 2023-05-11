package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
	"mime/multipart"
)

type IHistoricalService interface {
	Add(historyRow *dtos.HistoryCompleteDTO) *dtos.ValidationDTO
	AddHistoricalTransactions(file *multipart.FileHeader, userId int) *dtos.ValidationDTO
	GetAllHistorical(userId int) ([]*models.HistoricalModelWithOutJoins, *dtos.ValidationDTO)
	GetAmmountAtCreator(creatorId int) (uint64, *dtos.ValidationDTO)
	GetAfiliateHistorical(userId int, afiliateId int) (*dtos.HistoricalCompleteAfiliateDTO, *dtos.ValidationDTO)
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

func (service *historicalService) GetAllHistorical(userId int) ([]*models.HistoricalModelWithOutJoins, *dtos.ValidationDTO) {
	historicals, err := service.historyRepository.GetAll(userId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    23,
			Message: "An error ocurred when get all historical from user",
		}
	}

	return historicals, nil
}

func (service *historicalService) GetAmmountAtCreator(creatorId int) (uint64, *dtos.ValidationDTO) {
	receive, err := service.historyRepository.GetAmmountReceivedAtCreator(creatorId)
	if err != nil {
		return 0, &dtos.ValidationDTO{
			Code:    26,
			Message: "Error when search amount",
		}
	}

	paid, err := service.historyRepository.GetAmmountPaidAtCreator(creatorId)
	if err != nil {
		return 0, &dtos.ValidationDTO{
			Code:    27,
			Message: "Error when search amount",
		}
	}

	return (receive - paid), nil
}

func (service *historicalService) GetAfiliateHistorical(userId int, afiliateId int) (*dtos.HistoricalCompleteAfiliateDTO, *dtos.ValidationDTO) {
	creator, validationDTO := service.creatorService.GetCreator(userId)
	if validationDTO != nil {
		return nil, validationDTO
	}

	afiliateHistorical, err := service.historyRepository.GetHistoricalFromAfiliate(creator.CreatorId, afiliateId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    28,
			Message: "Error when search historical at afiliate",
		}
	}

	amountReceived, err := service.historyRepository.GetAmountReceivedFromAfiliate(creator.CreatorId, afiliateId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    29,
			Message: "Error when get amount at afiliate",
		}
	}

	historicalDTO := &dtos.HistoricalCompleteAfiliateDTO{
		AfiliateHistoricals: afiliateHistorical,
		Amount:              amountReceived,
	}
	return historicalDTO, nil
}
