package data_provider

import (
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
)

type PostgresProvider struct {}

func (p PostgresProvider) AddAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {
	panic("implement me")
}

func (p PostgresProvider) DeleteAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {
	panic("implement me")
}

func NewPostgresProvider() core.Database {
	return &PostgresProvider{}
}

