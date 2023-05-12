package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/exceptions"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type IAfiliatedService interface {
	AddAfiliate(inputAfiliate *dtos.AfiliatedDTO, userId int) *dtos.ValidationDTO
	FindAfiliate(name string, creatorId int) (*entitys.Afiliated, *dtos.ValidationDTO)
	GetAllAfiliates(userId int) ([]*models.AfiliateModel, *dtos.ValidationDTO)
}

type afiliatedService struct {
	afiliatedRepository repositorys.IAfliateRepository
}

func NewAfiliatedService() IAfiliatedService {
	return &afiliatedService{
		afiliatedRepository: repositorys.NewAfiliateRepository(),
	}
}

func (service *afiliatedService) AddAfiliate(inputAfiliate *dtos.AfiliatedDTO, userId int) *dtos.ValidationDTO {
	isCreated, err := service.afiliatedRepository.AddNewAfiliate(inputAfiliate, userId)

	if err != nil || !isCreated {
		log.Println(err)
		return &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeAddAfiliate,
			Message: exceptions.ErrorMessageddAfiliate,
		}
	}

	return nil
}

func (service *afiliatedService) FindAfiliate(name string, creatorId int) (*entitys.Afiliated, *dtos.ValidationDTO) {
	afiliate, err := service.afiliatedRepository.Find(name, creatorId)
	if err != nil || afiliate.ID == 0 {
		return nil, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeAfiliateNotFound,
			Message: exceptions.ErrorMessageAfiliateNotFound,
		}
	}

	return afiliate, nil
}

func (service *afiliatedService) GetAllAfiliates(userId int) ([]*models.AfiliateModel, *dtos.ValidationDTO) {
	afiliates, err := service.afiliatedRepository.GetAll(userId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeAllAfiliateNotFound,
			Message: exceptions.ErrorMessageAllAfiliateNotFound,
		}
	}

	return afiliates, nil
}
