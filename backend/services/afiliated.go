package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type IAfiliatedService interface {
	AddAfiliate(inputAfiliate *dtos.AfiliatedDTO, email string, idCreator int) *dtos.ValidationDTO
}

type afiliatedService struct {
	afiliatedRepository repositorys.IAfliateRepository
}

func NewAfiliatedService() IAfiliatedService {
	return &afiliatedService{
		afiliatedRepository: repositorys.NewAfiliateRepository(),
	}
}

func (service *afiliatedService) AddAfiliate(inputAfiliate *dtos.AfiliatedDTO, email string, idCreator int) *dtos.ValidationDTO {
	isCreated, err := service.afiliatedRepository.AddNewAfiliate(inputAfiliate, email, idCreator)

	if err != nil {
		log.Println(err)
		return &dtos.ValidationDTO{
			Code:    10,
			Message: "Error when add new afiliate",
		}
	} else if !isCreated {
		return &dtos.ValidationDTO{
			Code:    11,
			Message: "Invalid creator data to add afiliate",
		}
	}

	return nil
}
