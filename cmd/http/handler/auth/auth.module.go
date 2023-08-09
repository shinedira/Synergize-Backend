package auth

import (
	"synergize/backend-test/pkg/facades"

	"github.com/labstack/echo/v4"
)

type AuthModuleHandler struct{}

func (h *AuthModuleHandler) ApiV1() (*echo.Group, []any) {
	group := facades.Route.Group("api/v1/auth")

	return group, []any{
		&AuthPlayerHandler{},
	}
}
