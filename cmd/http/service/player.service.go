package service

import (
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/pkg/entity"
	"synergize/backend-test/pkg/facades"
)

type PlayerService struct{}

type PlayerServiceInterface interface {
	UpdateBalance(*models.Player, uint64) error
	Store(*models.Player) error
	FindAll(*models.PlayerFilterRequest) ([]models.Player, entity.Pagination, error)
	FindById(*models.Player) error
	FindByUsername(string) (*models.Player, error)
}

func (s *PlayerService) UpdateBalance(player *models.Player, amount uint64) error {
	player.Balance += amount

	if err := facades.DB.Save(&player).Error; err != nil {
		return err
	}

	return nil
}

func (s *PlayerService) Store(player *models.Player) error {
	if err := facades.DB.Create(player).Error; err != nil {
		return err
	}

	return nil
}

func (s *PlayerService) FindAll(req *models.PlayerFilterRequest) (players []models.Player, pagination entity.Pagination, err error) {
	bankIds := facades.DB.Scopes(
		req.FilterByBankName,
		req.FilterByBankAccountName,
		req.FilterByBankAccountNumber,
	).Model(&models.BankAccount{}).Select("player_id")

	if err := facades.DB.Scopes(
		entity.Paginate(&players, &pagination, &req.PaginationRequest, facades.DB),
		req.FilterByUsername,
		req.FilterByName,
	).
		Preload("Banks").
		Where("id IN (?)", bankIds).
		Find(&players).Error; err != nil {
		return players, pagination, err
	}

	return players, pagination, nil
}

func (s *PlayerService) FindById(player *models.Player) error {
	if err := facades.DB.Preload("Banks").Find(player).Error; err != nil {
		return err
	}

	return nil
}

func (s *PlayerService) FindByUsername(username string) (*models.Player, error) {
	var player models.Player

	if err := facades.DB.
		Where("username = ?", username).
		First(&player).Error; err != nil {
		return nil, err
	}

	return &player, nil
}
