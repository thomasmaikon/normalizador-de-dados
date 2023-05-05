package models

type User struct {
	ID   int    `gorm: "primaryKey; autoIncrement"`
	Name string `gorm: "not null"`

	LoginID int
	Login   Login `gorm: "foreignKey; constraint:OnUpdate:CASCADE"`
}
