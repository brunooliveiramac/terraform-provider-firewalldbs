package service

import (
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/data_provider"
)

type DatabaseProvider interface {
	AddIp(ipRule *entity.ServerFirewallIpRule, token string) error
	RemoveIp(ipRule *entity.ServerFirewallIpRule, token string) error
	SetNext(database DatabaseProvider)
	SetDBProvider(db core.Database)
}

func GetProvider() DatabaseProvider {

	noProvider := &NoProvider{}

	postgres := &Postgres{}
	azurePostgres := data_provider.NewAzurePostgresProvider()
	postgres.SetDBProvider(azurePostgres)
	postgres.SetNext(noProvider)

	mysql := &Mysql{}
	azureMysql := data_provider.NewAzureMysqlProvider()
	mysql.SetDBProvider(azureMysql)
	mysql.SetNext(postgres)

	return mysql
}

func contains(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}
