package bootstrap

import (
	"synergize/backend-test/cmd/providers/application"
	"synergize/backend-test/cmd/providers/authentication"
	"synergize/backend-test/cmd/providers/cache"
	"synergize/backend-test/cmd/providers/configuration"
	"synergize/backend-test/cmd/providers/database"
	"synergize/backend-test/cmd/providers/handler"
	"synergize/backend-test/pkg/entity"
)

var (
	//register config & provider here
	configs = []entity.ServiceProvider{
		&configuration.ConfigServiceProvider{},
		&application.AppServiceProvider{},
		&database.DatabaseServiceProvider{},
		&cache.CacheServiceProvider{},
		&handler.HandlerServiceProvider{},
		&authentication.AuthServiceProvider{},
	}
)

func Boot() {
	for _, config := range configs {
		config.Register()
		config.Boot()
	}
}
