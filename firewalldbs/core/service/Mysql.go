package service

import (
	"strings"
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
)

type Mysql struct {
	next DatabaseProvider
	db   core.Database
}

func (mysqlInstance *Mysql) AddIp(ipRule *entity.ServerFirewallIpRule, token string) error {
	ids := strings.Split(ipRule.ServerID, "/")
	if ipRule.ServerID == "" || contains(ids, "Microsoft.DBforMySQL") {
		err := mysqlInstance.db.AddAgentIp(ipRule, token)

		if err != nil {
			return err
		}

		return nil
	}

	return mysqlInstance.next.AddIp(ipRule, token)
}

func (mysqlInstance *Mysql) RemoveIp(ipRule *entity.ServerFirewallIpRule, token string) error {
	ids := strings.Split(ipRule.ServerID, "/")
	if ipRule.ServerID == "" || contains(ids, "Microsoft.DBforMySQL") {

		err := mysqlInstance.db.DeleteAgentIp(ipRule, token)

		if err != nil {
			return err
		}

		return nil
	}

	return mysqlInstance.next.RemoveIp(ipRule, token)
}

func (mysqlInstance *Mysql) SetNext(next DatabaseProvider) {
	mysqlInstance.next = next
}

func (mysqlInstance *Mysql) SetDBProvider(db core.Database) {
	mysqlInstance.db = db
}
