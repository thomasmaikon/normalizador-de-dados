package entitys

import "time"

type History struct {
	ID    int `gorm:"autoIncrement;primaryKey"`
	Value uint64
	Date  time.Time `gorm:"type:date"`

	CreatorID int
	Creator   Creator

	ProductID int
	Product   Product

	AfiliatedID int
	Afiliated   Afiliated

	TransactionID int
	Transaction   Transaction
}
