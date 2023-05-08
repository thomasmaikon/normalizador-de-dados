package repositorys

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/utils"

	"gorm.io/gorm"
)

type ICreatorRepository interface {
	CreateCreator(newCreator *dtos.CreatorDTO, userID int) error
}

type creatorRepository struct {
	db *gorm.DB
}

func NewCreatorRepository() ICreatorRepository {
	return &creatorRepository{
		db: utils.GetDB(),
	}
}

func (repository *creatorRepository) CreateCreator(newCreator *dtos.CreatorDTO, userID int) error {
	creator := entitys.Creator{
		Name:   newCreator.Name,
		UserID: userID,
	}

	err := repository.db.Create(&creator)

	return err.Error
}
