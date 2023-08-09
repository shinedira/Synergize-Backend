package database

import (
	"synergize/backend-test/cmd/models"
	"synergize/backend-test/pkg/facades"
)

type DatabaseServiceProvider struct{}

func (p *DatabaseServiceProvider) Boot() {
	facades.DB.AutoMigrate(
		models.Player{},
		models.BankAccount{},
		models.Transaction{},
	)
}

func (p *DatabaseServiceProvider) Register() {
	config := facades.Config
	config.Add("database", map[string]any{
		"pgsql": map[string]any{
			"host":     config.Env("DB_HOST", "127.0.0.1"),
			"port":     config.Env("DB_PORT", "3306"),
			"database": config.Env("DB_DATABASE", "backend-test"),
			"username": config.Env("DB_USERNAME", ""),
			"password": config.Env("DB_PASSWORD", ""),
		},
	})

	facades.DB = bootDatabase()
}
