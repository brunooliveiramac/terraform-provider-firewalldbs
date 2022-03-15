package service

import (
	"strings"
	"terraform-provider-firewalldbs/firewalldbs/core"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
)

type Postgres struct {
	next DatabaseProvider
	db core.Database
}

func (postgresInstance *Postgres) AddIp(ipRule *entity.ServerFirewallIpRule, token string)  error {
	ids := strings.Split(ipRule.ServerID, "/")
	if contains(ids, "Microsoft.DBforPostgreSQL") {
		err := postgresInstance.db.AddAgentIp(ipRule, token)

		if err != nil {
			return err
		}
	}
	return postgresInstance.next.AddIp(ipRule, token)
}

func (postgresInstance *Postgres) RemoveIp(ipRule *entity.ServerFirewallIpRule, token string)  error {
	ids := strings.Split(ipRule.ServerID, "/")
	if contains(ids, "Microsoft.DBforPostgreSQL") {
		err := postgresInstance.db.AddAgentIp(ipRule, token)

		if err != nil {
			return err
		}
	}
	return postgresInstance.next.RemoveIp(ipRule, token)
}


func (postgresInstance *Postgres) SetNext(next DatabaseProvider) {
	postgresInstance.next = next
}

func (postgresInstance *Postgres) SetDBProvider(db core.Database) {
	postgresInstance.db = db
}