package services

import (
	"bufio"
	"hubla/desafiofullstack/models"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type InormalizeDataService interface {
	GetNormalizedData(file *multipart.FileHeader) ([]*models.HistoryModel, error)
}

type normalizeDataService struct {
}

func NewNormalizeDataService() InormalizeDataService {
	return &normalizeDataService{}
}

func (normalize *normalizeDataService) GetNormalizedData(file *multipart.FileHeader) ([]*models.HistoryModel, error) {

	scanner, err := normalize.castingFile(file)
	if err != nil {
		return nil, err
	}

	historicals := make([]*models.HistoryModel, 0)

	for scanner.Scan() {
		line := scanner.Text()
		historcal := normalize.normalize(line)
		historicals = append(historicals, historcal)
	}

	return historicals, nil
}

func (normalize *normalizeDataService) castingFile(file *multipart.FileHeader) (*bufio.Scanner, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(src), err
}

func (normalize *normalizeDataService) normalize(inputData string) *models.HistoryModel {
	pattern := "([0-9])([T0-9:-]+)([A-Z- ]+)[ ]+([0-9]+)([A-Z ]+)"
	regex := regexp.MustCompile(pattern)
	submatches := regex.FindAllStringSubmatch(inputData, -1)

	layoutTime := "2006-01-02T15:04:05-07:00"

	idTransaction, err := strconv.Atoi(submatches[0][1])
	if err != nil {
		return nil
	}

	date, err := time.Parse(layoutTime, submatches[0][2])
	if err != nil {
		return nil
	}

	description := strings.TrimSpace(submatches[0][3])

	LeftOver, err := strconv.ParseUint(submatches[0][4], 10, 64)
	if err != nil {
		return nil
	}

	afiliate := submatches[0][5]

	return &models.HistoryModel{
		IdTransactionType:  idTransaction,
		Date:               date,
		ProductDescription: description,
		Value:              LeftOver,
		Afiliate:           afiliate,
	}
}
