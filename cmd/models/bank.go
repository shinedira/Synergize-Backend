package models

import (
	"time"

	"gorm.io/gorm"
)

// base model
type (
	BankAccount struct {
		ID            uint64     `gorm:"primaryKey" json:"id"`
		PlayerId      uint64     `gorm:"index; not null" json:"player_id"`
		BankName      string     `gorm:"type:varchar(255); not null" json:"bank_name"`
		AccountName   string     `gorm:"type:varchar(255); not null" json:"account_name"`
		AccountNumber uint64     `gorm:"not null" json:"account_number"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     time.Time  `json:"updated_at"`
		DeletedAt     *time.Time `json:"-"`

		Player Player `gorm:"foreignKey:PlayerId"`
	}
)

// request
type (
	BankAccountRequest struct {
		BankName      string `json:"bank_name" validate:"required,max=255"`
		AccountName   string `json:"account_name" validate:"required"`
		AccountNumber uint64 `json:"account_number" validate:"required,numeric"`
	}

	BankAccountFilterRequest struct {
		FilterBankName          string `query:"filters[bank][name]"`
		FilterBankAccountName   string `query:"filters[bank][account_name]"`
		FilterBankAccountNumber string `query:"filters[bank][account_number]"`
	}
)

// response
type (
	BankAccountResponse struct {
		ID            uint64 `json:"id"`
		BankName      string `json:"name"`
		AccountName   string `json:"account_name"`
		AccountNumber uint64 `json:"account_number"`
	}
)

func (m *BankAccountFilterRequest) FilterByBankName(db *gorm.DB) *gorm.DB {
	if m.FilterBankName == "" {
		return db
	}

	return db.Where("LOWER(bank_name) = LOWER(?)", m.FilterBankName)
}

func (m *BankAccountFilterRequest) FilterByBankAccountName(db *gorm.DB) *gorm.DB {
	if m.FilterBankAccountName == "" {
		return db
	}

	return db.Where("LOWER(account_name) = LOWER(?)", m.FilterBankAccountName)
}

func (m *BankAccountFilterRequest) FilterByBankAccountNumber(db *gorm.DB) *gorm.DB {
	if m.FilterBankAccountNumber == "" {
		return db
	}

	return db.Where("account_number = (?)", m.FilterBankAccountNumber)
}
