package repositorys

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/utils"

	"gorm.io/gorm"
)

type ILoginRepository interface {
	Create(inputLogin *dtos.LoginDTO) (*entitys.Login, error)
	Validate(inputLogin *dtos.LoginDTO) (*entitys.Login, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository() ILoginRepository {
	return &loginRepository{db: utils.GetDB()}
}

func (repository *loginRepository) Create(input *dtos.LoginDTO) (*entitys.Login, error) {
	newLogin := &entitys.Login{
		Email:    input.Email,
		Password: input.Password,
	}

	err := repository.db.Create(newLogin)

	return newLogin, err.Error
}

func (loginRepository *loginRepository) Validate(input *dtos.LoginDTO) (*entitys.Login, error) {
	var login entitys.Login
	err := loginRepository.db.Table("logins").Where("email = ? AND password = ?", input.Email, input.Password).Scan(&login)

	return &login, err.Error
}
