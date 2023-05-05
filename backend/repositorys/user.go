package repositorys

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/utils"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(inputUser dtos.CreateUseDTO, loginId int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &userRepository{db: utils.GetDB()}
}

func (repository *userRepository) CreateUser(inputUser dtos.CreateUseDTO, loginId int) error {

	newUser := &models.User{
		Name:    inputUser.Name,
		LoginID: loginId,
	}

	err := repository.db.Create(newUser)
	return err.Error
}
