package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"
)

type IHistoryRepository interface {
	AddHistoryRow(history *dtos.HistoryCompleteDTO) (bool, error)
	Begin()
	Commit()
	Rollback()
}

type historyRepository struct {
	uow *utils.UnitOfWork
}

func NewHistoryService() IHistoryRepository {
	return &historyRepository{
		uow: utils.GetUnitOfWork(),
	}
}

func (repository *historyRepository) AddHistoryRow(history *dtos.HistoryCompleteDTO) (bool, error) {
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

func (repository *historyRepository) Begin() {
	repository.uow.Begin()
}

func (repository *historyRepository) Commit() {
	repository.uow.Commit()
}

func (repository *historyRepository) Rollback() {
	repository.uow.Rollback()
}
