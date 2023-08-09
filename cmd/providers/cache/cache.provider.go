package cache

import (
	"synergize/backend-test/pkg/facades"
)

type CacheServiceProvider struct{}

const CACHE_PREFIX = "synergize_cache"

func (p *CacheServiceProvider) Boot() {}

func (p *CacheServiceProvider) Register() {
	config := facades.Config
	config.Add("database", map[string]any{
		"redis": map[string]any{
			"host":     config.Env("REDIS_HOST", "127.0.0.1"),
			"password": config.Env("REDIS_PASSWORD", ""),
			"port":     config.Env("REDIS_PORT", 6379),
			"database": config.Env("REDIS_DB", 0),
		},
	})

	facades.Cache = bootCache()
}
