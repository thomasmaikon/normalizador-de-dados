package models

type User struct {
	ID   int    `gorm: "primaryKey; autoIncrement"`
	Name string `gorm: "not null"`

	LoginID int
	Login   Login `gorm: "constraint:OnUpdate:CASCADE"`
}

func (user User) TableName() string {
	return "user"
}
