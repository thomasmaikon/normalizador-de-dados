package dtos

import (
	"hubla/desafiofullstack/models"
	"time"
)

type HistoryCompleteDTO struct {
	IdCreator         int
	Date              time.Time
	Value             uint64
	IdProduct         int
	IdAfiliated       int
	IdTransactionType int
}

type HistoricalCompleteAfiliateDTO struct {
	AfiliateHistoricals *[]models.HistoricalModelWithOutJoins
	Amount              uint64
}
