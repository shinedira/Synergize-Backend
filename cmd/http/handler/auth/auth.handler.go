package auth

import (
	"net/http"
	"synergize/backend-test/cmd/http/middleware"
	"synergize/backend-test/cmd/http/service"
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/pkg/facades"
	"synergize/backend-test/pkg/helper"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthPlayerHandler struct {
	playerService service.PlayerServiceInterface
	authService   service.AuthPlayerServiceInterface
}

func (h *AuthPlayerHandler) Boot(route *echo.Group) {
	h.playerService = &service.PlayerService{}
	h.authService = &service.AuthPlayerService{}

	route.POST("/register", h.register)
	route.POST("/login", h.login)
	route.POST("/refresh", h.refresh)
	route.POST("/logout", h.logout, middleware.AuthencateMiddleware)
}

func (h *AuthPlayerHandler) register(c echo.Context) error {
	var player models.Player

	if err := helper.MapAndValidate(c, &player, &models.RegisterRequest{}); err != nil {
		return err
	}

	if err := h.playerService.Store(&player); err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.Resp{"message": "Pendaftaran berhasil"})
}

func (h *AuthPlayerHandler) login(c echo.Context) error {
	var request models.LoginRequest

	if err := helper.Validation(c, &request); err != nil {
		return err
	}

	admin, err := h.playerService.FindByUsername(request.Username)
	if err != nil {
		return helper.UnauthenticatedErrorReponse()
	}

	err = helper.HashCheck(request.Password, admin.Password)
	if err != nil {
		return helper.UnauthenticatedErrorReponse()
	}

	accessToken, err := facades.Auth.CreateAccessToken(admin.Username)
	if err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	if err = h.authService.StoreAccessTokenSession(models.AuthStoreSession{
		ID:        accessToken.ID,
		Token:     accessToken.Token,
		ExpiredAt: accessToken.ExpiredAt,
		Username:  request.Username,
		UserAgent: c.Request().UserAgent(),
		ClientIP:  c.RealIP(),
	}); err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	refreshToken, err := facades.Auth.CreateRefreshToken(admin.Username)
	if err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	if err = h.authService.StoreRefreshTokenSession(models.AuthStoreSession{
		ID:        refreshToken.ID,
		Token:     refreshToken.Token,
		ExpiredAt: refreshToken.ExpiredAt,
		Username:  request.Username,
		UserAgent: c.Request().UserAgent(),
		ClientIP:  c.RealIP(),
	}); err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.LoginResponse{
		AccessToken:           accessToken.Token,
		AccessTokenExpiredAt:  accessToken.ExpiredAt,
		RefreshToken:          refreshToken.Token,
		RefreshTokenExpiredAt: refreshToken.ExpiredAt,
		User:                  admin,
	})
}

func (h *AuthPlayerHandler) refresh(c echo.Context) error {
	var request models.AuthRefreshRequest

	if err := helper.Validation(c, &request); err != nil {
		return err
	}

	refreshTokenPayload, err := facades.Auth.VerifyToken(request.RefreshToken)
	if err != nil {
		return helper.UnauthenticatedErrorReponse()
	}

	session, err := h.authService.FindOneRefreshSession(request.RefreshToken)
	if err != nil {
		return helper.UnauthenticatedErrorReponse()
	}

	if session.Revoked {
		return helper.UnauthenticatedErrorReponse("session revoked")
	}

	if session.Username != refreshTokenPayload.Username {
		return helper.UnauthenticatedErrorReponse("incorect session user")
	}

	if session.Token != request.RefreshToken {
		return helper.UnauthenticatedErrorReponse("mismatched session token")
	}

	if time.Now().After(session.ExpiredAt) {
		return helper.UnauthenticatedErrorReponse("expired token")
	}

	accessToken, err := facades.Auth.CreateAccessToken(refreshTokenPayload.Username)
	if err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	if err = h.authService.StoreAccessTokenSession(models.AuthStoreSession{
		ID:        accessToken.ID,
		Token:     accessToken.Token,
		ExpiredAt: accessToken.ExpiredAt,
		Username:  session.Username,
		UserAgent: c.Request().UserAgent(),
		ClientIP:  c.RealIP(),
	}); err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models.AuthRefreshResponse{
		AccessToken:          accessToken.Token,
		AccessTokenExpiredAt: accessToken.ExpiredAt,
	})
}

func (h *AuthPlayerHandler) logout(c echo.Context) error {
	token := c.Get("token").(string)

	if err := h.authService.Logout(token); err != nil {
		return helper.ErrorReponse(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.Resp{
		"message": "logout berhasil",
	})
}
