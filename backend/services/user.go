package services

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/exceptions"
	"hubla/desafiofullstack/repositorys"
	"log"
)

type IUserService interface {
	CreateUser(input dtos.CreateUseDTO) *dtos.ValidationDTO
	FindUser(email string) (*entitys.User, *dtos.ValidationDTO)
}

type userService struct {
	ILoginService
	repositorys.IUserRepository
}

func NewUserService() IUserService {
	return &userService{
		IUserRepository: repositorys.NewUserRepository(),
		ILoginService:   NewLoginService(),
	}
}

func (service *userService) CreateUser(input dtos.CreateUseDTO) *dtos.ValidationDTO {
	login, err := service.ILoginService.CreateLogin(input.Login)

	if err != nil {
		return err
	}

	input.Login.Email = login.Email
	input.Login.Password = login.Password

	erroUser := service.IUserRepository.CreateUser(input, login.ID)

	if erroUser != nil {
		log.Fatal(erroUser.Error())
		return &dtos.ValidationDTO{
			Code:    exceptions.ErrorCodeCreateUser,
			Message: exceptions.ErrorMessageCreateUser,
		}
	}

	return nil
}

func (service *userService) FindUser(email string) (*entitys.User, *dtos.ValidationDTO) {
	user, err := service.IUserRepository.FindUser(email)
	if err != nil {
		return nil, &dtos.ValidationDTO{
			Code:    5,
			Message: "User doesn't find",
		}
	}
	return user, nil
}
