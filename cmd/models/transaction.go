package models

import (
	"time"
)

type TransactionType string

const (
	Expense TransactionType = "expense"
	TopUp   TransactionType = "topup"
)

type (
	Transaction struct {
		ID        uint64          `gorm:"primaryKey" json:"id"`
		PlayerId  uint64          `gorm:"index; not null" json:"player_id"`
		Amount    uint64          `gorm:"not null" json:"amount"`
		Type      TransactionType `gorm:"type:varchar(255); not null" json:"type"`
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *time.Time      `json:"-"`

		Player Player `gorm:"foreignKey:PlayerId"`
	}
)

type (
	TopUpRequest struct {
		Amount uint64 `validate:"numeric"`
	}
)

type (
	TopUpResponse struct {
		ID     uint64          `json:"id"`
		Amount uint64          `json:"amount"`
		Type   TransactionType `json:"type"`
		Player PlayerResponse  `json:"player"`
	}
)
