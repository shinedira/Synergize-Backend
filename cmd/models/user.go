package models

import (
	"synergize/backend-test/pkg/entity"
	"synergize/backend-test/pkg/helper"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// base model
type (
	Player struct {
		ID        uint64     `gorm:"primaryKey" json:"id" param:"id"`
		Name      string     `gorm:"type:varchar(255); not null" json:"name"`
		Balance   uint64     `gorm:"not null" json:"balance"`
		Username  string     `gorm:"unique; type:varchar(255); not null" json:"username"`
		Password  string     `gorm:"type:varchar(255); not null" json:"-"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"-"`

		Banks        []BankAccount `gorm:"foreignKey:PlayerId"`
		Transactions []Transaction `gorm:"foreignKey:PlayerId"`
	}

	AuthStoreSession struct {
		ID        uuid.UUID `json:"id"`
		Token     string    `json:"token"`
		Revoked   bool      `json:"revoked"`
		ExpiredAt time.Time `json:"expired_at"`
		Username  string    `json:"username"`
		UserAgent string    `json:"user_agent"`
		ClientIP  string    `json:"client_ip"`
	}
)

// request
type (
	RegisterRequest struct {
		Name     string `validate:"required,max=255"`
		Username string `validate:"required,alphanum,max=255,unique=players:username"`
		Password string `validate:"required,alphanum,max=255"`
	}

	LoginRequest struct {
		Username string `validate:"required,max=255"`
		Password string `validate:"alphanum"`
	}

	AuthRefreshRequest struct {
		RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
	}

	PlayerFilterRequest struct {
		entity.PaginationRequest
		BankAccountFilterRequest
		FilterUsername string `query:"filters[username]"`
		FilterName     string `query:"filters[name]"`
	}
)

// response
type (
	LoginResponse struct {
		AccessToken           string    `json:"access_token"`
		AccessTokenExpiredAt  time.Time `json:"access_token_expired_at"`
		RefreshToken          string    `json:"refresh_token"`
		RefreshTokenExpiredAt time.Time `json:"refresh_token_expired_at"`
		User                  any       `json:"user"`
	}

	AuthRefreshResponse struct {
		AccessToken          string    `json:"access_token"`
		AccessTokenExpiredAt time.Time `json:"access_token_expired_at"`
	}

	PlayerResponse struct {
		ID       uint64 `json:"id"`
		Name     string `json:"name"`
		Balance  uint64 `json:"balance"`
		Username string `json:"username"`
	}

	PlayerDetailResponse struct {
		ID       uint64                `json:"id"`
		Name     string                `json:"name"`
		Balance  uint64                `json:"balance"`
		Username string                `json:"username"`
		Banks    []BankAccountResponse `json:"bank_accounts"`
	}
)

func (u *Player) BeforeCreate(tx *gorm.DB) (err error) {
	u.Balance = 0

	if u.Password != "" {
		u.Password, err = helper.Hash(u.Password)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *PlayerFilterRequest) FilterByUsername(db *gorm.DB) *gorm.DB {
	if m.FilterUsername == "" {
		return db
	}

	return db.Where("LOWER(username) = LOWER(?)", m.FilterUsername)
}

func (m *PlayerFilterRequest) FilterByName(db *gorm.DB) *gorm.DB {
	if m.FilterName == "" {
		return db
	}

	return db.Where("name ILIKE ?", "%"+m.FilterName+"%")
}
