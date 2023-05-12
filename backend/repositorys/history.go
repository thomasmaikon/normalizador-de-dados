package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"
)

type IHistoricalRepository interface {
	AddHistoryRow(history *dtos.HistoryCompleteDTO) (bool, error)
	GetAll(userId int) ([]*models.HistoricalModelWithOutJoins, error)
	GetAmmountReceivedAtCreator(creatorId int) (uint64, error)
	GetAmmountPaidAtCreator(creatorId int) (uint64, error)
	GetHistoricalFromAfiliate(creatorId int, afiliateId int) (*[]models.HistoricalModelWithOutJoins, error)
	GetAmountReceivedFromAfiliate(creatorId int, afiliateId int) (uint64, error)
	Begin()
	Commit()
	Rollback()
}

type historicalRepository struct {
	uow *utils.UnitOfWork
}

func NewHistoricalService() IHistoricalRepository {
	return &historicalRepository{
		uow: utils.GetUnitOfWork(),
	}
}

func (repository *historicalRepository) AddHistoryRow(history *dtos.HistoryCompleteDTO) (bool, error) {
	result := repository.uow.GetDB().Exec(
		querys.AddingHistoryRow,
		sql.Named(querys.NamedDate, history.Date),
		sql.Named(querys.NamedValue, history.Value),
		sql.Named(querys.NamedCreatorsId, history.IdCreator),
		sql.Named(querys.NamedProductId, history.IdProduct),
		sql.Named(querys.NamedAfiliatedId, history.IdAfiliated),
		sql.Named(querys.NamedTransactionId, history.IdTransactionType),
	)

	return result.RowsAffected == 1, result.Error
}

func (repository *historicalRepository) GetAll(userId int) ([]*models.HistoricalModelWithOutJoins, error) {

	var historicals []*models.HistoricalModelWithOutJoins

	result := repository.uow.GetDB().Raw(
		querys.GetAllDataFromUser,
		sql.Named(querys.NamedUserId, userId),
	).Scan(&historicals)

	return historicals, result.Error
}

func (repository *historicalRepository) GetAmmountReceivedAtCreator(creatorId int) (uint64, error) {
	var received uint64
	err := repository.uow.GetDB().Raw(
		querys.GetAmmountReceivedValueAtCreator,
		sql.Named(querys.NamedID, creatorId),
	).Scan(&received)

	return received, err.Error
}

func (repository *historicalRepository) GetAmmountPaidAtCreator(creatorId int) (uint64, error) {
	var received uint64
	err := repository.uow.GetDB().Raw(
		querys.GetAmmountPaidValueAtCreator,
		sql.Named(querys.NamedID, creatorId),
	).Scan(&received)

	return received, err.Error
}

func (repository *historicalRepository) GetHistoricalFromAfiliate(creatorId int, afiliateId int) (*[]models.HistoricalModelWithOutJoins, error) {
	var historicals *[]models.HistoricalModelWithOutJoins

	result := repository.uow.GetDB().Raw(
		querys.GetAllDataFromAfiliate,
		sql.Named(querys.NamedCreatorsId, creatorId),
		sql.Named(querys.NamedAfiliatedId, afiliateId),
	).Scan(&historicals)

	return historicals, result.Error
}

func (repository *historicalRepository) GetAmountReceivedFromAfiliate(creatorId int, afiliateId int) (uint64, error) {
	var received uint64

	err := repository.uow.GetDB().Raw(
		querys.GetAmountReceivedFromAfiliate,
		sql.Named(querys.NamedCreatorsId, creatorId),
		sql.Named(querys.NamedAfiliatedId, afiliateId),
	).Scan(&received)

	return received, err.Error
}

func (repository *historicalRepository) Begin() {
	repository.uow.Begin()
}

func (repository *historicalRepository) Commit() {
	repository.uow.Commit()
}

func (repository *historicalRepository) Rollback() {
	repository.uow.Rollback()
}
