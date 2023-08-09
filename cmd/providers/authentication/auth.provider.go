package authentication

import "synergize/backend-test/pkg/facades"

type AuthServiceProvider struct{}

const (
	JWT_SCREATE            = "7WS74j2hdi44P5rGzvoTpbRLhOG5whiF"
	ACCESS_TOKEN_DURATION  = 4320
	REFRESH_TOKEN_DURATION = 262800
	AUTH_HEADER            = "Authorization"
	AUTH_TYPE              = "bearer"
	AUTH_PAYLOAD_KEY       = "auth_payload"
)

func (p *AuthServiceProvider) Boot() {}

func (p *AuthServiceProvider) Register() {
	facades.Auth = NewJWTService(JWT_SCREATE)
}
