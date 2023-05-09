package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"
)

type IAfliateRepository interface {
	AddNewAfiliate(inputAfiliate *dtos.AfiliatedDTO, userId int) (bool, error)
	Find(name string, creatorId int) (*entitys.Afiliated, error)
}

type afiliateRepository struct {
	uow *utils.UnitOfWork
}

func NewAfiliateRepository() IAfliateRepository {
	return &afiliateRepository{
		uow: utils.GetUnitOfWork(),
	}
}

func (repository *afiliateRepository) AddNewAfiliate(inputAfiliate *dtos.AfiliatedDTO, userId int) (bool, error) {

	result := repository.uow.GetDB().Exec(querys.AddAfiliate,
		sql.Named(querys.NamedName, inputAfiliate.Name),
		sql.Named(querys.NamedUserId, userId),
	)

	return result.RowsAffected > 0, result.Error
}

func (repository *afiliateRepository) Find(name string, creatorId int) (*entitys.Afiliated, error) {
	var afialte entitys.Afiliated
	err := repository.uow.GetDB().Table("afiliateds").Select("*").Where("name = ? AND creator_id = ?", name, creatorId).Scan(&afialte)

	return &afialte, err.Error
}
