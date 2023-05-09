package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/repositorys"
)

type IUserService interface {
	CreateUser(input dtos.CreateUseDTO) (*models.UserModel, *dtos.ValidationDTO)
	FindUser(email string) (*entitys.User, *dtos.ValidationDTO)
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

func (service *userService) CreateUser(input dtos.CreateUseDTO) (*models.UserModel, *dtos.ValidationDTO) {

	service.userRepository.Begin()

	newUser, err := service.userRepository.CreateUser(input)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    17,
			Message: "Doesn`t possible create user",
		}
	}

	validationDTO := service.loginService.CreateLogin(input.Login, newUser.ID)
	if validationDTO != nil {
		service.userRepository.RollBack()
		return nil, validationDTO
	}

	service.userRepository.Commit()

	return &models.UserModel{
		UserId: newUser.ID,
	}, nil
}

func (service *userService) FindUser(email string) (*entitys.User, *dtos.ValidationDTO) {
	user, err := service.userRepository.FindUser(email)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    5,
			Message: "User doesn't find",
		}
	}
	return user, nil
}
