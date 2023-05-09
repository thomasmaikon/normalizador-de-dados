package repositorys

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/utils"
)

type ICreatorRepository interface {
	CreateCreator(newCreator *dtos.CreatorDTO, userID int) error
	Find(userId int) (*entitys.Creator, error)
}

type creatorRepository struct {
	uow *utils.UnitOfWork
}

func NewCreatorRepository() ICreatorRepository {
	return &creatorRepository{
		uow: utils.GetUnitOfWork(),
	}
}

func (repository *creatorRepository) CreateCreator(newCreator *dtos.CreatorDTO, userID int) error {
	creator := entitys.Creator{
		Name:   newCreator.Name,
		UserID: userID,
	}

	err := repository.uow.GetDB().Create(&creator)

	return err.Error
}

func (repository *creatorRepository) Find(userId int) (*entitys.Creator, error) {
	var creator entitys.Creator
	err := repository.uow.GetDB().Table("creators").Where("USER_ID = ?", userId).Scan(&creator)

	return &creator, err.Error
}
