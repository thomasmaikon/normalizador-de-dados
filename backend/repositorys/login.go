package repositorys

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/utils"

	"gorm.io/gorm"
)

type ILoginRepository interface {
	Create(inputLogin *dtos.LoginDTO) (*models.Login, error)
	Validate(inputLogin *dtos.LoginDTO) error
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository() ILoginRepository {
	return &loginRepository{db: utils.GetDB()}
}

func (repository *loginRepository) Create(input *dtos.LoginDTO) (*models.Login, error) {
	newLogin := &models.Login{
		Email:    input.Email,
		Password: input.Password,
	}

	err := repository.db.Create(newLogin)

	return newLogin, err.Error
}

func (loginRepository *loginRepository) Validate(input *dtos.LoginDTO) error {

	err := loginRepository.db.Find(&models.Login{Email: input.Email, Password: input.Password})

	return err.Error
}
