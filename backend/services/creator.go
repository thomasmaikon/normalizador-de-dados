package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
	"log"
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
		log.Println(err.Error())
		return &dtos.ValidationDTO{
			Code:    6,
			Message: "Does not possible create an creator",
		}
	}

	return nil
}

func (service *creatorService) GetCreator(userId int) (*models.CreatorModel, *dtos.ValidationDTO) {
	creator, err := service.creatorRepository.Find(userId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    20,
			Message: "Creator doesn`t finded",
		}
	}

	return &models.CreatorModel{
		CreatorId: creator.ID,
		Name:      creator.Name,
		LeftOver:  0,
	}, nil
}
