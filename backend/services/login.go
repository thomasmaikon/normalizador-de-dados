package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type ILoginService interface {
	CreateLogin(inputLogin dtos.LoginDTO) (*models.Login, *dtos.ValidationDTO)
	ValidateCredential(inputLogin dtos.LoginDTO) *dtos.ValidationDTO
}

type loginService struct {
	repositorys.ILoginRepository
}

func NewLoginService() ILoginService {
	return &loginService{
		ILoginRepository: repositorys.NewLoginRepository(),
	}
}

func (service *loginService) CreateLogin(inputLogin dtos.LoginDTO) (*models.Login, *dtos.ValidationDTO) {
	login, err := service.ILoginRepository.Create(&inputLogin)
	if err != nil {
		log.Println(err.Error())
		return nil, &dtos.ValidationDTO{
			Code:    2,
			Message: "Faild to create login",
		}
	}

	return login, nil
}

func (service *loginService) ValidateCredential(inputLogin dtos.LoginDTO) *dtos.ValidationDTO {
	err := service.ILoginRepository.Validate(&inputLogin)
	if err != nil {
		return &dtos.ValidationDTO{
			Code:    3,
			Message: "User not found",
		}
	}

	return nil
}
