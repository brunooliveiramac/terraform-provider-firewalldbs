package service

import (
	"errors"
	"fmt"
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
)

type NoProvider struct {
	next DatabaseProvider
	db   core.Database
}

func (no *NoProvider) AddIp(ipRule *entity.ServerFirewallIpRule, token string) error {
	msg := fmt.Sprintf("No Provider Match for server %s server_id %s!", ipRule.ServerName, ipRule.ServerID)
	return errors.New(msg)
}

func (no *NoProvider) RemoveIp(ipRule *entity.ServerFirewallIpRule, token string) error {
	msg := fmt.Sprintf("No Provider Match for server %s server_id %s!", ipRule.ServerName, ipRule.ServerID)
	return errors.New(msg)
}

func (no *NoProvider) SetNext(next DatabaseProvider) {
	no.next = next
}

func (no *NoProvider) SetDBProvider(db core.Database) {
	no.db = db
}
