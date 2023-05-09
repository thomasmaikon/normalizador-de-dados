package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"
)

type IUserRepository interface {
	CreateUser(inputUser dtos.CreateUseDTO) (*entitys.User, error)
	FindUser(email string) (*entitys.User, error)
	Begin()
	Commit()
	RollBack()
}

type userRepository struct {
	uow *utils.UnitOfWork
}

func NewUserRepository() IUserRepository {
	return &userRepository{uow: utils.GetUnitOfWork()}
}

func (repository *userRepository) CreateUser(inputUser dtos.CreateUseDTO) (*entitys.User, error) {

	newUser := &entitys.User{
		Name: inputUser.Name,
	}

	err := repository.uow.GetDB().Create(newUser)
	return newUser, err.Error
}
func (repository *userRepository) FindUser(email string) (*entitys.User, error) {
	var user entitys.User

	err := repository.uow.GetDB().Raw(
		querys.FinUserByEmail,
		sql.Named(querys.NamedEmail, email),
	).Scan(&user)

	return &user, err.Error
}

func (repository *userRepository) Begin() {
	repository.uow.Begin()
}
func (repository *userRepository) Commit() {
	repository.uow.Commit()
}
func (repository *userRepository) RollBack() {
	repository.uow.Rollback()
}
