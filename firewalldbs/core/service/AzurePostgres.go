package service

import (
	"strings"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/data_provider"
)

type AzurePostgres struct {
	next DatabaseProvider
}

func (postgresInstance *AzurePostgres) AddIp(ipRule *entity.ServerFirewallIpRule, token string)  error {
	ids := strings.Split(ipRule.ServerID, "/")
	if contains(ids, "Microsoft.DBforPostgreSQL") {
		postgresProvider := data_provider.NewPostgresProvider()
		err := postgresProvider.AddAgentIp(ipRule, token)

		if err != nil {
			return err
		}
	}
	return postgresInstance.next.AddIp(ipRule, token)
}

func (postgresInstance *AzurePostgres) RemoveIp(ipRule *entity.ServerFirewallIpRule, token string)  error {
	ids := strings.Split(ipRule.ServerID, "/")
	if contains(ids, "Microsoft.DBforPostgreSQL") {
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
