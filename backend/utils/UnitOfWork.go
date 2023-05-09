package utils

import "gorm.io/gorm"

type UnitOfWork struct {
	db *gorm.DB
}

func (uow *UnitOfWork) GetDB() *gorm.DB {
	return uow.db
}

func (uow *UnitOfWork) Begin() {
	uow.db = uow.db.Begin()
}

func (uow *UnitOfWork) Commit() {
	uow.db.Commit()
}

func (uow *UnitOfWork) Rollback() {
	uow.db.Rollback()
}
