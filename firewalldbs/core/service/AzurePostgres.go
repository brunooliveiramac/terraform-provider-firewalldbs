package service

import (
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/data_provider"
)

type AzurePostgres struct {
	next DatabaseProvider
}

func (postgresInstance *AzurePostgres) AddIp(ipRule *entity.ServerFirewallIpRule, token string)  error {
	if ipRule.ServerID == "Postgres" {
		postgresProvider := data_provider.NewPostgresProvider()
		err := postgresProvider.AddAgentIp(ipRule, token)

		if err != nil {
			return err
		}
	}
	return postgresInstance.next.AddIp(ipRule, token)
}

func (postgresInstance *AzurePostgres) RemoveIp(ipRule *entity.ServerFirewallIpRule, token string)  error {
	if ipRule.ServerID == "Postgres" {
		postgresProvider := data_provider.NewPostgresProvider()
		err := postgresProvider.AddAgentIp(ipRule, token)

		if err != nil {
			return err
		}
	}
	return postgresInstance.next.RemoveIp(ipRule, token)
}


func (postgresInstance *AzurePostgres) SetNext(next DatabaseProvider) {
	postgresInstance.next = next
}
