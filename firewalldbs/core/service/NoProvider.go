package service

import (
	"errors"
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
)

type NoProvider struct {
	next DatabaseProvider
	db core.Database
}

func (no *NoProvider) AddIp(ipRule *entity.ServerFirewallIpRule, token string)  error {
	return errors.New("No Provider Match!")
}

func (no *NoProvider) RemoveIp(ipRule *entity.ServerFirewallIpRule, token string)  error {
	return errors.New("No Provider Match!")
}

func (no *NoProvider) SetNext(next DatabaseProvider) {
	no.next = next
}

func (no *NoProvider) SetDBProvider(db core.Database) {
	no.db = db
}
