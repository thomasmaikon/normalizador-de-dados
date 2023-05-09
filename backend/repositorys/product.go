package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateNewProduct(newProduct *dtos.ProductDTO, email string, idCreator int) (bool, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository() IProductRepository {
	return &productRepository{
		db: utils.GetDB(),
	}
}

func (repository *productRepository) CreateNewProduct(newProduct *dtos.ProductDTO, email string, idCreator int) (bool, error) {

	err := repository.db.Exec(
		querys.CreateNewProduct,
		sql.Named(querys.NamedDescription, newProduct.Description),
		sql.Named(querys.NamedCreatorsId, idCreator),
		sql.Named(querys.NamedEmail, email),
		sql.Named(querys.NamedPrice, newProduct.Price),
	)
	
	return err.RowsAffected > 0, err.Error
}
