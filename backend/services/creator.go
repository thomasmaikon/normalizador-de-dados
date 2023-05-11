package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/exceptions"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
)

type ICreatorService interface {
	CreateNewCreator(newCreator *dtos.CreatorDTO, userID int) *dtos.ValidationDTO
	GetCreator(userId int) (*models.CreatorModel, *dtos.ValidationDTO)
}

type creatorService struct {
	creatorRepository repositorys.ICreatorRepository
}

func NewCreatorSerivce() ICreatorService {
	return &creatorService{
		creatorRepository: repositorys.NewCreatorRepository(),
	}
}

func (service *creatorService) CreateNewCreator(newCreator *dtos.CreatorDTO, userID int) *dtos.ValidationDTO {

	err := service.creatorRepository.CreateCreator(newCreator, userID)

	if err != nil {
		return &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeFaildCreateCreator,
			Message: exceptions.ErrorMessageFaildCreateCreator,
		}
	}

	return nil
}

func (service *creatorService) GetCreator(userId int) (*models.CreatorModel, *dtos.ValidationDTO) {
	creator, err := service.creatorRepository.Find(userId)
	if err != nil || creator.ID == 0 {
		return nil, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeFaildCreateCreator,
			Message: exceptions.ErrorMessageFaildCreatorNotFound,
		}
	}

	return &models.CreatorModel{
		CreatorId: creator.ID,
		Name:      creator.Name,
	}, nil
}
