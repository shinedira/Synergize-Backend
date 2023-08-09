package user

import (
	"net/http"
	"synergize/backend-test/cmd/http/service"
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/pkg/entity"
	"synergize/backend-test/pkg/helper"

	"github.com/labstack/echo/v4"
)

type PlayerHandler struct {
	playerService service.PlayerServiceInterface
}

func (h *PlayerHandler) Boot(route *echo.Group) {
	h.playerService = &service.PlayerService{}

	route.GET("/user", h.findAll)
	route.GET("/user/:id", h.findOne)
}

func (h *PlayerHandler) findAll(c echo.Context) error {
	var request models.PlayerFilterRequest

	if err := c.Bind(&request); err != nil {
		return err
	}

	data, pagination, err := h.playerService.FindAll(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var players []models.PlayerResponse

	helper.Clone(&players, &data)

	return c.JSON(http.StatusOK, &entity.PaginationResponse{
		Data:       players,
		Pagination: pagination,
	})
}

func (h *PlayerHandler) findOne(c echo.Context) error {
	var request models.Player

	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := h.playerService.FindById(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var player models.PlayerDetailResponse

	helper.Clone(&player, &request)

	return c.JSON(http.StatusOK, player)
}
