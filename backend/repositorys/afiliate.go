package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"
)

type IAfliateRepository interface {
	AddNewAfiliate(inputAfiliate *dtos.AfiliatedDTO, email string, idCreator int) (bool, error)
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

func (repository *afiliateRepository) AddNewAfiliate(inputAfiliate *dtos.AfiliatedDTO, email string, idCreator int) (bool, error) {

	result := repository.uow.GetDB().Exec(querys.AddAfiliate,
		sql.Named(querys.NamedName, inputAfiliate.Name),
		sql.Named(querys.NamedCreatorsId, idCreator),
		sql.Named(querys.NamedEmail, email),
	)

	return result.RowsAffected > 0, result.Error
}

func (repository *afiliateRepository) Find(name string, creatorId int) (*entitys.Afiliated, error) {
	var afialte entitys.Afiliated
	err := repository.uow.GetDB().Table("AFILIATEDS").Select("*").Where("name = '?' AND creator_id = ?", name, creatorId).Scan(&afialte)

	return &afialte, err.Error
}
