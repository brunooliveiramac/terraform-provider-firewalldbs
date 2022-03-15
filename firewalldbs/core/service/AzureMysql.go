package service

import (
	"strings"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/data_provider"
)

type AzureMysql struct {
	next DatabaseProvider
}

func (mysqlInstance *AzureMysql) AddIp(ipRule *entity.ServerFirewallIpRule, token string) error {
	ids := strings.Split(ipRule.ServerID, "/")
	if contains(ids, "Microsoft.DBforMySQL") {
		mysqlProvider := data_provider.NewMysqlProvider()
		err := mysqlProvider.AddAgentIp(ipRule, token)

		if err != nil {
			return err
		}

		return nil
	}

	return mysqlInstance.next.AddIp(ipRule, token)
}

func (mysqlInstance *AzureMysql) RemoveIp(ipRule *entity.ServerFirewallIpRule, token string) error {
	ids := strings.Split(ipRule.ServerID, "/")
	if contains(ids, "Microsoft.DBforMySQL") {
		mysqlProvider := data_provider.NewMysqlProvider()
		err := mysqlProvider.DeleteAgentIp(ipRule, token)

		if err != nil {
			return err
		}

		return nil
	}

	return mysqlInstance.next.RemoveIp(ipRule, token)
}

func (mysqlInstance *AzureMysql) SetNext(next DatabaseProvider) {
	mysqlInstance.next = next
}
