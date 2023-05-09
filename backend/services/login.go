package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
)

type ILoginService interface {
	CreateLogin(inputLogin dtos.LoginDTO, userId int) *dtos.ValidationDTO
	ValidateCredential(inputLogin dtos.LoginDTO) (*models.UserModel, *dtos.ValidationDTO)
}

type loginService struct {
	loginRepository repositorys.ILoginRepository
}

func NewLoginService() ILoginService {
	return &loginService{
		loginRepository: repositorys.NewLoginRepository(),
	}
}

func (service *loginService) CreateLogin(inputLogin dtos.LoginDTO, userId int) *dtos.ValidationDTO {
	err := service.loginRepository.Create(&inputLogin, userId)
	if err != nil {
		return &dtos.ValidationDTO{
			Code:    2,
			Message: "Faild to create login",
		}
	}

	return nil
}

func (service *loginService) ValidateCredential(inputLogin dtos.LoginDTO) (*models.UserModel, *dtos.ValidationDTO) {
	login, err := service.loginRepository.Validate(&inputLogin)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    3,
			Message: "User not found",
		}
	}

	return &models.UserModel{UserId: login.UserID}, nil
}
