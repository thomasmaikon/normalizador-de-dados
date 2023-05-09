package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"

	"gorm.io/gorm"
)

type IAfliateRepository interface {
	AddNewAfiliate(inputAfiliate *dtos.AfiliatedDTO, email string, idCreator int) (bool, error)
}

type afiliateRepository struct {
	db *gorm.DB
}

func NewAfiliateRepository() IAfliateRepository {
	return &afiliateRepository{
		db: utils.GetDB(),
	}
}

func (repository *afiliateRepository) AddNewAfiliate(inputAfiliate *dtos.AfiliatedDTO, email string, idCreator int) (bool, error) {

	result := repository.db.Exec(querys.AddAfiliate,
		sql.Named(querys.NamedName, inputAfiliate.Name),
		sql.Named(querys.NamedCreatorsId, idCreator),
		sql.Named(querys.NamedEmail, email),
	)

	return result.RowsAffected > 0, result.Error
}
