package application

import (
	"synergize/backend-test/pkg/facades"
	"time"
)

type AppServiceProvider struct{}

func (p *AppServiceProvider) Boot() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
}

func (p *AppServiceProvider) Register() {
	config := facades.Config
	config.Add("app", map[string]any{
		"name": config.Env("APP_NAME", "BackendTest"),
		"env":  config.Env("APP_ENV", "development"),
		"host": config.Env("APP_HOST", "127.0.0.1:3000"),
	})
}
