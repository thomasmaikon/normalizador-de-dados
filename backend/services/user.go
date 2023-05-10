package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
)

type IUserService interface {
	CreateUser(input dtos.UserDTO) (*models.UserModel, *dtos.ValidationDTO)
	GetUser(inputLogin *dtos.LoginDTO) (*models.UserModel, *dtos.ValidationDTO)
}

type userService struct {
	loginService   ILoginService
	userRepository repositorys.IUserRepository
}

func NewUserService() IUserService {
	return &userService{
		userRepository: repositorys.NewUserRepository(),
		loginService:   NewLoginService(),
	}
}

func (service *userService) CreateUser(input dtos.UserDTO) (*models.UserModel, *dtos.ValidationDTO) {

	login, validationDTO := service.loginService.CreateLogin(input.Login)
	if validationDTO != nil {
		return nil, validationDTO
	}

	newUser, err := service.userRepository.CreateUser(input, login.LoginId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    17,
			Message: "Doesn`t possible create user",
		}
	}

	return &models.UserModel{
		UserId: newUser.ID,
	}, nil
}

func (service *userService) GetUser(inputLogin *dtos.LoginDTO) (*models.UserModel, *dtos.ValidationDTO) {
	loginModel, validationDto := service.loginService.ValidateCredential(inputLogin)
	if validationDto != nil {
		return nil, validationDto
	}

	user, err := service.userRepository.FindUser(loginModel.LoginId)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    19,
			Message: "User not find",
		}
	}
	return &models.UserModel{UserId: user.ID}, nil
}
