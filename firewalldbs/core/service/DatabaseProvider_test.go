package service

import (
	"github.com/stretchr/testify/assert"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/core/service/mock"
	"testing"
)

func TestShouldApplyMySqlProviderWhenIdContainsMySql(t *testing.T) {
	//Given
	mockAzureMysql := mock.MockAzureMysqlProvider{}

	mockAzureMysql.On("AddAgentIp").Return(nil)
	mockAzureMysql.On("DeleteAgentIp").Return(nil)

	serverInfo := entity.ServerFirewallIpRule{
		ServerID:      "/subscriptions/5c92b4a1-d813-42e0-804d-0c0e64218b27/resourceGroups/bees-eu-sbx-brunoxyy/providers/Microsoft.DBforMySQL/servers/brunoxyy-6fu-north-eu-sandbox",
		IP:            "192.168.0.1",
		ServerName:    "some-server",
		ResourceGroup: "some-resource-group",
		Subscription:  "some-subscription",
	}

	token := "SomeToken"

	mysql := &Mysql{}
	mysql.SetDBProvider(&mockAzureMysql)
	mysql.SetNext(nil)

	//When
	err := mysql.AddIp(&serverInfo, token)

	//Then
	//Behavioral assertion
	assert.Equal(t, nil, err)
}

