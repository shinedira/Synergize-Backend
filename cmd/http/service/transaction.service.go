package service

import (
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/pkg/facades"
)

type TransactionService struct{}

type TransactionServiceInterface interface {
	Store(*models.Transaction) error
}

func (s *TransactionService) Store(transaction *models.Transaction) error {
	if err := facades.DB.Save(transaction).Error; err != nil {
		return err
	}

	return nil
}
