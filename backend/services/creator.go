package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type ICreatorService interface {
	CreateNewCreator(newCreator *dtos.CreatorDTO, userID int) *dtos.ValidationDTO
}

type creatorService struct {
	IUserService
	repositorys.ICreatorRepository
}

func NewCreatorSerivce() ICreatorService {
	return &creatorService{
		IUserService:       NewUserService(),
		ICreatorRepository: repositorys.NewCreatorRepository(),
	}
}

func (service *creatorService) CreateNewCreator(newCreator *dtos.CreatorDTO, userID int) *dtos.ValidationDTO {

	err := service.ICreatorRepository.CreateCreator(newCreator, userID)

	if err != nil {
		log.Println(err.Error())
		return &dtos.ValidationDTO{
			Code:    6,
			Message: "Does not possible create an creator",
		}
	}

	return nil
}
