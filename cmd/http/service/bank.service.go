package service

import (
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/pkg/facades"
)

type BankAccountService struct{}

type BankAccountServiceInterface interface {
	Store(*models.BankAccount) error
}

func (s *BankAccountService) Store(bankAccount *models.BankAccount) error {
	if err := facades.DB.Save(bankAccount).Error; err != nil {
		return err
	}

	return nil
}
