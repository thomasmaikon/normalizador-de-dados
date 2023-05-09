package models

import "time"

type HistoryModel struct {
	IdTransactionType  int
	Date               time.Time
	ProductDescription string
	Value              uint64
	Afiliate           string
}
