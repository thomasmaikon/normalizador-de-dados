package repositorys

import (
	"database/sql"
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/entitys"
	"hubla/desafiofullstack/querys"
	"hubla/desafiofullstack/utils"
)

type IProductRepository interface {
	CreateNewProduct(newProduct *dtos.ProductDTO, email string, idCreator int) (bool, error)
	Find(description string, creatorId int) (*entitys.Product, error)
}

type productRepository struct {
	uow *utils.UnitOfWork
}

func NewProductRepository() IProductRepository {
	return &productRepository{
		uow: utils.GetUnitOfWork(),
	}
}

func (repository *productRepository) CreateNewProduct(newProduct *dtos.ProductDTO, email string, idCreator int) (bool, error) {

	err := repository.uow.GetDB().Exec(
		querys.CreateNewProduct,
		sql.Named(querys.NamedDescription, newProduct.Description),
		sql.Named(querys.NamedCreatorsId, idCreator),
		sql.Named(querys.NamedEmail, email),
		sql.Named(querys.NamedPrice, newProduct.Price),
	)

	return err.RowsAffected > 0, err.Error
}

func (repository *productRepository) Find(description string, creatorId int) (*entitys.Product, error) {

	var product entitys.Product

	err := repository.uow.GetDB().
		Table("products").
		Select("*").
		Where("creator_id = ? AND description = ?", creatorId, description).
		Scan(&product)

	return &product, err.Error
}
