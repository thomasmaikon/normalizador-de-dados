package repositorys

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/utils"

	"gorm.io/gorm"
)

type ILoginRepository interface {
	Create(inputLogin *dtos.LoginDTO, userId int) error
	Validate(inputLogin *dtos.LoginDTO) (*entitys.Login, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository() ILoginRepository {
	return &loginRepository{db: utils.GetDB()}
}

func (repository *loginRepository) Create(input *dtos.LoginDTO, userId int) error {
	newLogin := &entitys.Login{
		Email:    input.Email,
		Password: input.Password,
		UserID:   userId,
	}

	err := repository.db.Create(newLogin)

	return err.Error
}

func (loginRepository *loginRepository) Validate(input *dtos.LoginDTO) (*entitys.Login, error) {
	var login entitys.Login
	err := loginRepository.db.Find(&entitys.Login{Email: input.Email, Password: input.Password}).Scan(&login)

	return &login, err.Error
}
