package bank

import (
	"net/http"
	"synergize/backend-test/cmd/http/service"
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/cmd/providers/authentication"
	"synergize/backend-test/pkg/entity"
	"synergize/backend-test/pkg/helper"

	"github.com/labstack/echo/v4"
)

type BankAccountHandler struct {
	bankAccountService service.BankAccountServiceInterface
	playerService      service.PlayerServiceInterface
}

func (h *BankAccountHandler) Boot(route *echo.Group) {
	h.bankAccountService = &service.BankAccountService{}
	h.playerService = &service.PlayerService{}

	route.POST("/bank-account", h.store)
}

func (h *BankAccountHandler) store(c echo.Context) error {
	auth := c.Get(authentication.AUTH_PAYLOAD_KEY).(*entity.AuthPayload)
	player, err := h.playerService.FindByUsername(auth.Username)
	if err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	bankAccount := &models.BankAccount{
		Player: *player,
	}

	if err := helper.MapAndValidate(c, &bankAccount, &models.BankAccountRequest{}); err != nil {
		return err
	}

	if err := h.bankAccountService.Store(bankAccount); err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.Resp{
		"message": "Akun bank player berhasil di tambahkan",
		"data": models.BankAccountResponse{
			ID:            bankAccount.ID,
			BankName:      bankAccount.BankName,
			AccountName:   bankAccount.AccountName,
			AccountNumber: bankAccount.AccountNumber,
		},
	})
}
