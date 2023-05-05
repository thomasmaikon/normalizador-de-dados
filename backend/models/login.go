package models

type Login struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
