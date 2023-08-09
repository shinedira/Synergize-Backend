package entity

import (
	"time"
)

type ServiceProvider interface {
	Boot()
	Register()
}

type ConfigBoot interface {
	Boot()
}

type Config interface {
	Env(envName string, defaultValue ...any) any
	Add(name string, configuration any)
	Get(path string, defaultValue ...any) any
	GetString(path string, defaultValue ...any) string
	GetInt(path string, defaultValue ...any) int
	GetBool(path string, defaultValue ...any) bool
}

type Cache interface {
	Push(string, any) error
	Retrieve(string) []string
	Remove(string, int)
	Pop(string) []string
	Get(string, any) any
	Has(string) bool
	Set(string, any, time.Duration) error
	Pull(string, any) any
	Add(string, any, time.Duration) bool
	Remember(string, time.Duration, func() any) (any, error)
	RememberForever(string, func() any) (any, error)
	Forever(string, any) bool
	Forget(string) bool
	Flush() bool
}

type JwtServiceInterface interface {
	CreateToken(string, time.Duration) (string, AuthPayload, error)
	VerifyToken(string) (*AuthPayload, error)
	CreateAccessToken(username string) (*AuthTokenResult, error)
	CreateRefreshToken(username string) (*AuthTokenResult, error)
}
