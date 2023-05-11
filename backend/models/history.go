package models

import "time"

type HistoryModel struct {
	IdTransactionType  int
	Date               time.Time
	ProductDescription string
	Value              uint64
	Afiliate           string
}

type HistoricalModelWithOutJoins struct {
	AfiliateName           string `gorm:"column:afiliate"`
	ProductDescription     string `gorm:"column:product"`
	TransactionDescription string `gorm:"column:transaction"`
	Value                  uint64
	Date                   time.Time
}
