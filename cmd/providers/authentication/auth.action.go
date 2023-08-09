package authentication

import (
	"synergize/backend-test/pkg/entity"
	"time"
)

func (j *JwtService) CreateAccessToken(username string) (*entity.AuthTokenResult, error) {
	duration := ACCESS_TOKEN_DURATION * time.Minute

	token, payload, err := j.CreateToken(username, duration)
	if err != nil {
		return nil, err
	}

	return &entity.AuthTokenResult{
		Token:       token,
		AuthPayload: payload,
	}, nil
}

func (j *JwtService) CreateRefreshToken(username string) (*entity.AuthTokenResult, error) {
	duration := REFRESH_TOKEN_DURATION * time.Minute

	token, payload, err := j.CreateToken(username, duration)
	if err != nil {
		return nil, err
	}

	return &entity.AuthTokenResult{
		Token:       token,
		AuthPayload: payload,
	}, nil
}
