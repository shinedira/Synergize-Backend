package facades

import (
	"synergize/backend-test/pkg/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	// Instance is the global instance of the application.
	Config entity.Config
	DB     *gorm.DB
	Cache  entity.Cache
	Route  *echo.Echo
	Auth   entity.JwtServiceInterface
)
