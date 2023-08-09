package configuration

import "synergize/backend-test/pkg/facades"

type ConfigServiceProvider struct{}

func (p *ConfigServiceProvider) Boot() {}

func (p *ConfigServiceProvider) Register() {
	facades.Config = BootConfig()
}
