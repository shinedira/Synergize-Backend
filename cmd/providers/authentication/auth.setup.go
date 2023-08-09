package authentication

import (
	"errors"
	"synergize/backend-test/pkg/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrUnauthorized = errors.New("unauthenticated")
	ErrExpiredToken = errors.New("token has been expired")
	ErrInvalidToken = errors.New("invalid token")
	TokenType       = "bearer"
)

type JwtService struct {
	secreatKey string
}

func NewJWTService(secreatKey string) entity.JwtServiceInterface {
	return &JwtService{secreatKey}
}

func (j *JwtService) CreateToken(username string, duration time.Duration) (string, entity.AuthPayload, error) {
	payload, err := NewPayload(username, duration)
	maskPayload := entity.AuthPayload(*payload)
	// maskPayload := (*entity.AuthPayload)(unsafe.Pointer(payload))

	if err != nil {
		return "", maskPayload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(j.secreatKey))

	return token, maskPayload, err
}

func (j *JwtService) VerifyToken(token string) (*entity.AuthPayload, error) {
	Func := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return []byte(j.secreatKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, Func)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if errors.Is(v.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return &entity.AuthPayload{
		ID:        payload.ID,
		Username:  payload.Username,
		IssuedAt:  payload.IssuedAt,
		ExpiredAt: payload.ExpiredAt,
	}, nil
}
