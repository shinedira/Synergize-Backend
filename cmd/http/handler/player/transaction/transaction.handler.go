package transaction

import (
	"net/http"
	"synergize/backend-test/cmd/http/service"
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/cmd/providers/authentication"
	"synergize/backend-test/pkg/entity"
	"synergize/backend-test/pkg/helper"

	"github.com/labstack/echo/v4"
)

type TransactionPlayerHandler struct {
	transactionService service.TransactionServiceInterface
	playerService      service.PlayerServiceInterface
}

func (h *TransactionPlayerHandler) Boot(route *echo.Group) {
	h.transactionService = &service.TransactionService{}
	h.playerService = &service.PlayerService{}

	g := route.Group("/transaction")
	g.POST("/topup", h.store)
}

func (h *TransactionPlayerHandler) store(c echo.Context) error {
	//do some security before add data here
	//

	//store history transaction
	auth := c.Get(authentication.AUTH_PAYLOAD_KEY).(*entity.AuthPayload)
	player, err := h.playerService.FindByUsername(auth.Username)
	if err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	var request models.TopUpRequest
	transaction := &models.Transaction{
		Player: *player,
		Type:   models.TopUp,
	}

	if err := helper.MapAndValidate(c, &transaction, &request); err != nil {
		return err
	}

	if err := h.transactionService.Store(transaction); err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	//update balance here
	if err := h.playerService.UpdateBalance(player, request.Amount); err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.Resp{
		"message": "Topup berhasil",
		"data": models.TopUpResponse{
			ID:     transaction.ID,
			Amount: transaction.Amount,
			Type:   transaction.Type,
			Player: models.PlayerResponse{
				ID:       player.ID,
				Username: player.Username,
				Name:     player.Name,
				Balance:  player.Balance,
			},
		},
	})
}
