package core

import "terraform-provider-firewalldbs/firewalldbs/core/entity"

type Server struct {
	ID string
}

type Database interface {
	AddAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error)
	DeleteAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error)
}
