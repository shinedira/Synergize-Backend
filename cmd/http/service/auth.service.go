package service

import (
	"encoding/json"
	"fmt"
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/pkg/facades"
)

type AuthPlayerService struct{}

type AuthPlayerServiceInterface interface {
	StoreAccessTokenSession(models.AuthStoreSession) error
	StoreRefreshTokenSession(models.AuthStoreSession) error
	Logout(string) error
	FindOneRefreshSession(string) (models.AuthStoreSession, error)
	IsValidSession(string, string) bool
}

const (
	ACCESS_TOKEN_KEY  = "access_token"
	REFRESH_TOKEN_KEY = "refresh_token"
)

func (s *AuthPlayerService) StoreAccessTokenSession(session models.AuthStoreSession) error {
	if err := storeSession(session, ACCESS_TOKEN_KEY); err != nil {
		return err
	}

	return nil
}

func (s *AuthPlayerService) StoreRefreshTokenSession(session models.AuthStoreSession) error {
	if err := storeSession(session, REFRESH_TOKEN_KEY); err != nil {
		return err
	}

	return nil
}

func (s *AuthPlayerService) Logout(token string) error {
	tokens := facades.Cache.Retrieve(ACCESS_TOKEN_KEY)
	data := unmarshalSession(tokens)

	_, idx, ok := findSessionByToken(data, token)
	if ok {
		facades.Cache.Remove(ACCESS_TOKEN_KEY, idx)
		return nil
	}

	return fmt.Errorf("logout session not found")
}

func (s *AuthPlayerService) FindOneRefreshSession(token string) (models.AuthStoreSession, error) {
	session, _, ok := findOneSession(token, REFRESH_TOKEN_KEY)

	if ok {
		return session, nil
	}

	return models.AuthStoreSession{}, fmt.Errorf("refresh session not found")
}

func (s *AuthPlayerService) IsValidSession(token string, t string) bool {
	tokens := facades.Cache.Retrieve(t)
	data := unmarshalSession(tokens)

	session, _, ok := findSessionByToken(data, token)
	if ok {
		return !session.Revoked
	}

	return false
}

func storeSession(payload models.AuthStoreSession, t string) error {
	if err := facades.Cache.Push(t, payload); err != nil {
		return err
	}

	return nil
}

func findOneSession(token string, t string) (models.AuthStoreSession, int, bool) {
	tokens := facades.Cache.Retrieve(t)
	data := unmarshalSession(tokens)

	return findSessionByToken(data, token)
}

func findSessionByToken(data []models.AuthStoreSession, token string) (models.AuthStoreSession, int, bool) {
	for idx, t := range data {
		if t.Token == token {
			return t, idx, true
		}
	}

	return models.AuthStoreSession{}, -1, false
}

func unmarshalSession(data []string) []models.AuthStoreSession {
	var mapData []models.AuthStoreSession

	for _, itemJSON := range data {
		var itemData models.AuthStoreSession
		if err := json.Unmarshal([]byte(itemJSON), &itemData); err != nil {
			fmt.Println("Error decoding JSON:", err)
		}
		mapData = append(mapData, itemData)
	}

	return mapData
}
