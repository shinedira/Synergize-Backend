package player

import (
	"synergize/backend-test/cmd/http/handler/player/bank"
	"synergize/backend-test/cmd/http/handler/player/transaction"
	"synergize/backend-test/cmd/http/handler/player/user"
	"synergize/backend-test/cmd/http/middleware"
	"synergize/backend-test/pkg/facades"

	"github.com/labstack/echo/v4"
)

type PlayerModuleHandler struct{}

func (h *PlayerModuleHandler) ApiV1() (*echo.Group, []any) {
	group := facades.Route.Group("api/v1/player", middleware.AuthencateMiddleware)

	return group, []any{
		&bank.BankAccountHandler{},
		&transaction.TransactionPlayerHandler{},
		&user.PlayerHandler{},
	}
}
