package handler

import (
	"synergize/backend-test/cmd/http/handler/auth"
	"synergize/backend-test/cmd/http/handler/player"
	"synergize/backend-test/pkg/facades"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HandlerServiceProvider struct{}

func (p *HandlerServiceProvider) Boot() {
	validation := &CustomValidator{validator: validator.New()}

	facades.Route.Validator = validation
	bootCustomeValidation(validation.validator)
}

func (p *HandlerServiceProvider) Register() {
	facades.Route = echo.New()

	bootHandler([]any{
		&auth.AuthModuleHandler{},
		&player.PlayerModuleHandler{},
		// other modules here
	})
}
