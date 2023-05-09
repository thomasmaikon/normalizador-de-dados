package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"
)

type IUserRepository interface {
	CreateUser(inputUser dtos.CreateUseDTO, loginId int) (*entitys.User, error)
	FindUser(loginId int) (*entitys.User, error)
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

func (repository *userRepository) CreateUser(inputUser dtos.CreateUseDTO, loginId int) (*entitys.User, error) {

	newUser := &entitys.User{
		Name:    inputUser.Name,
		LoginID: loginId,
	}

	err := repository.uow.GetDB().Create(newUser)
	return newUser, err.Error
}
func (repository *userRepository) FindUser(loginId int) (*entitys.User, error) {
	var user entitys.User

	err := repository.uow.GetDB().Raw(
		querys.FinUserByLoginId,
		sql.Named(querys.NamedID, loginId),
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
