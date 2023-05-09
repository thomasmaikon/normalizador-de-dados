package dtos

import "time"

type HistoryCompleteDTO struct {
	IdCreator         int
	Date              time.Time
	Value             uint64
	IdProduct         int
	IdAfiliated       int
	IdTransactionType int
}
