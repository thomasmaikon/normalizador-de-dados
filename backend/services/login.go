package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/exceptions"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
)

type ILoginService interface {
	CreateLogin(inputLogin dtos.LoginDTO) (*models.LoginModel, *dtos.ValidationDTO)
	ValidateCredential(inputLogin *dtos.LoginDTO) (*models.LoginModel, *dtos.ValidationDTO)
}

type loginService struct {
	loginRepository repositorys.ILoginRepository
}

func NewLoginService() ILoginService {
	return &loginService{
		loginRepository: repositorys.NewLoginRepository(),
	}
}

func (service *loginService) CreateLogin(inputLogin dtos.LoginDTO) (*models.LoginModel, *dtos.ValidationDTO) {
	login, err := service.loginRepository.Create(&inputLogin)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeFaildCreateLogin,
			Message: exceptions.ErrorMessageFaildCreateLogin,
		}
	}

	return &models.LoginModel{LoginId: login.ID}, nil
}

func (service *loginService) ValidateCredential(inputLogin *dtos.LoginDTO) (*models.LoginModel, *dtos.ValidationDTO) {
	login, err := service.loginRepository.Validate(inputLogin)
	if err != nil || login.ID == 0{
		return nil, &dtos.ValidationDTO{
			Code:   exceptions.ErrorCodeFaildLogin ,
			Message: exceptions.ErrorMessageFaildLogin,
		}
	}

	return &models.LoginModel{LoginId: login.ID}, nil
}
