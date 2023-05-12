package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/exceptions"
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
	if err != nil || !isComplete {
		return &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeFaildAddTransaction,
			Message: exceptions.ErrorMessageFaildAddTransaction,
		}
	}

	return nil
}

func (service *historicalService) AddHistoricalTransactions(file *multipart.FileHeader, userId int) *dtos.ValidationDTO {
	normalizedData, err := service.normalizedDataService.GetNormalizedData(file)
	if err != nil {
		return &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeFaildNormalizeFile,
			Message: exceptions.ErrorMessageFaildNormalizeFile,
		}
	}

	creator, validationDTO := service.creatorService.GetCreator(userId)
	if validationDTO != nil {
		return validationDTO
	}

	service.historyRepository.Begin()
	for _, historcal := range normalizedData {

		afiliate, validationDTO := service.afiliatedService.FindAfiliate(historcal.Afiliate, creator.CreatorId)
		if validationDTO != nil {
			service.historyRepository.Rollback()
			return validationDTO
		}

		product, validationDTO := service.productService.FindProduct(historcal.ProductDescription, creator.CreatorId)
		if validationDTO != nil {
			service.historyRepository.Rollback()
			return validationDTO
		}

		validationDTO = service.Add(&dtos.HistoryCompleteDTO{
			Date:              historcal.Date,
			IdCreator:         creator.CreatorId,
			Value:             historcal.Value,
			IdProduct:         product.ID,
			IdAfiliated:       afiliate.ID,
			IdTransactionType: historcal.IdTransactionType,
		})

		if validationDTO != nil {
			service.historyRepository.Rollback()
			return validationDTO
		}
	}
	service.historyRepository.Commit()
	return nil
}

func (service *historicalService) GetAllHistorical(userId int) ([]*models.HistoricalModelWithOutJoins, *dtos.ValidationDTO) {
	historicals, err := service.historyRepository.GetAll(userId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeHistorical,
			Message: exceptions.ErrorMessageHistorical,
		}
	}

	return historicals, nil
}

func (service *historicalService) GetAmmountAtCreator(creatorId int) (uint64, *dtos.ValidationDTO) {
	receive, err := service.historyRepository.GetAmmountReceivedAtCreator(creatorId)
	if err != nil {
		return 0, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeAmountFromCreator,
			Message: exceptions.ErrorMessageAmountFromCreator,
		}
	}

	paid, err := service.historyRepository.GetAmmountPaidAtCreator(creatorId)
	if err != nil {
		return 0, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeAmountFromCreator,
			Message: exceptions.ErrorMessageAmountFromCreator,
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
			Code:    exceptions.ErrorCodeNotFoundHistoricalFromAfiliate,
			Message: exceptions.ErrorMessageNotFoundHistoricalFromAfiliate,
		}
	}

	amountReceived, err := service.historyRepository.GetAmountReceivedFromAfiliate(creator.CreatorId, afiliateId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeAmountFromAfiliate,
			Message: exceptions.ErrorMessageAmountFromAfiliate,
		}
	}

	historicalDTO := &dtos.HistoricalCompleteAfiliateDTO{
		AfiliateHistoricals: afiliateHistorical,
		Amount:              amountReceived,
	}
	return historicalDTO, nil
}
