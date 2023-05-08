package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(inputUser dtos.CreateUseDTO, loginId int) error
	FindUser(email string) (*entitys.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &userRepository{db: utils.GetDB()}
}

func (repository *userRepository) CreateUser(inputUser dtos.CreateUseDTO, loginId int) error {

	newUser := &entitys.User{
		Name:    inputUser.Name,
		LoginID: loginId,
	}

	err := repository.db.Create(newUser)
	return err.Error
}
func (repository *userRepository) FindUser(email string) (*entitys.User, error) {
	var user entitys.User

	err := repository.db.Raw(
		querys.FinUserByEmail,
		sql.Named(querys.NamedEmail, email),
	).Scan(&user)

	return &user, err.Error
}
