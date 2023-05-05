package models

type Login struct {
	ID       int    `gorm: "primaryKey; autoIncrement"`
	Email    string `gorm: "not null; unique"`
	Password string `gorm: "not null"`
}
