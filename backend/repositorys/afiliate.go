package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"
)

type IAfliateRepository interface {
	AddNewAfiliate(inputAfiliate *dtos.AfiliatedDTO, userId int) (bool, error)
	Find(name string, creatorId int) (*entitys.Afiliated, error)
	GetAll(userId int) ([]*models.AfiliateModel, error)
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

	result := repository.uow.GetDB().Exec(
		querys.AddAfiliate,
		sql.Named(querys.NamedName, inputAfiliate.Name),
		sql.Named(querys.NamedUserId, userId),
	)

	return result.RowsAffected == 1, result.Error
}

func (repository *afiliateRepository) Find(name string, creatorId int) (*entitys.Afiliated, error) {
	var afialte entitys.Afiliated
	err := repository.uow.GetDB().Table("afiliateds").Select("*").Where("name = ? AND creator_id = ?", name, creatorId).Scan(&afialte)

	return &afialte, err.Error
}

func (repository *afiliateRepository) GetAll(userId int) ([]*models.AfiliateModel, error) {
	var afiliates []*models.AfiliateModel

	result := repository.uow.GetDB().Raw(
		querys.GetAllAfiliates,
		sql.Named(querys.NamedUserId, userId),
	).Scan(&afiliates)

	return afiliates, result.Error
}
