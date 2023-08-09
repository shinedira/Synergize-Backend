package middleware

import (
	"fmt"
	"strings"
	"synergize/backend-test/cmd/http/service"
	"synergize/backend-test/cmd/providers/authentication"
	"synergize/backend-test/pkg/facades"
	"synergize/backend-test/pkg/helper"

	"github.com/labstack/echo/v4"
)

var (
	authHeader     string = "Authorization"
	authType       string = "bearer"
	authPayloadKey string = "auth_payload"
)

func AuthencateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(authHeader)

		if len(header) == 0 {
			return helper.UnauthenticatedErrorReponse("authorization header is not provided")
		}

		fields := strings.Fields(header)
		if len(fields) < 2 {
			return helper.UnauthenticatedErrorReponse("invalid authorization header format")
		}

		getAuthType := strings.ToLower(fields[0])
		if getAuthType != authType {
			return helper.UnauthenticatedErrorReponse(fmt.Errorf("unsupported authorization type %s", authType).Error())
		}

		accessToken := fields[1]
		authService := &service.AuthPlayerService{}
		if !authService.IsValidSession(accessToken, service.ACCESS_TOKEN_KEY) {
			return helper.UnauthenticatedErrorReponse(authentication.ErrInvalidToken.Error())
		}

		payload, err := facades.Auth.VerifyToken(accessToken)
		if err != nil {
			return helper.UnauthenticatedErrorReponse(err.Error())
		}

		c.Set(authPayloadKey, payload)
		c.Set("token", accessToken)

		return next(c)
	}
}
