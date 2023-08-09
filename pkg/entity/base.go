package entity

import (
	"time"

	"github.com/google/uuid"
)

type AuthPayload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type AuthTokenResult struct {
	Token string
	AuthPayload
}
