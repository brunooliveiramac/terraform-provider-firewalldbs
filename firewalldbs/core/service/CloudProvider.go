package service

import (
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
)

type DatabaseProvider interface {
	AddIp(ipRule *entity.ServerFirewallIpRule, token string) error
	RemoveIp(ipRule *entity.ServerFirewallIpRule, token string) error
	SetNext(database DatabaseProvider)
}

func GetProvider () DatabaseProvider {
	noProvider := &NoProvider{}
	postgres := &AzurePostgres{}
	postgres.SetNext(noProvider)
	mysql := &AzureMysql{}
	mysql.SetNext(postgres)
	return mysql
}
