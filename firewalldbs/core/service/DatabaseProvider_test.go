package service

import (
	"github.com/stretchr/testify/assert"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
	"terraform-provider-firewalldbs/firewalldbs/core/service/mock"
	"testing"
)

func TestShouldApplyMySqlProviderWhenIdContainsMySql(t *testing.T) {
	//Given
	mockProvider := mock.MockProvider{}

	mockProvider.On("AddAgentIp").Return(nil)
	mockProvider.On("DeleteAgentIp").Return(nil)

	serverInfo := entity.ServerFirewallIpRule{
		ServerID:      "/subscriptions/5c92b4a1-d813-42e0-804d-0c0e64218b27/resourceGroups/bees-eu-sbx-brunoxyy/providers/Microsoft.DBforMySQL/servers/brunoxyy-6fu-north-eu-sandbox",
		IP:            "192.168.0.1",
		ServerName:    "some-server",
		ResourceGroup: "some-resource-group",
		Subscription:  "some-subscription",
	}

	token := "SomeToken"

	postgres := &Postgres{}
	postgres.SetDBProvider(&mockProvider)
	postgres.SetNext(nil)

	mysql := &Mysql{}
	mysql.SetDBProvider(&mockProvider)
	mysql.SetNext(postgres)

	//When
	err := mysql.AddIp(&serverInfo, token)

	//Then
	assert.Equal(t, nil, err)
}

func TestShouldApplyPostgresProviderWhenIdContainsPostgres(t *testing.T) {
	//Given
	mockProvider := mock.MockProvider{}

	mockProvider.On("AddAgentIp").Return(nil)
	mockProvider.On("DeleteAgentIp").Return(nil)

	serverInfo := entity.ServerFirewallIpRule{
		ServerID:      "/subscriptions/5c92b4a1-d813-42e0-804d-0c0e64218b27/resourceGroups/bees-eu-sbx-brunoxyy/providers/Microsoft.DBforPostgreSQL/servers/brunoxyy-6fu-north-eu-sandbox",
		IP:            "192.168.0.1",
		ServerName:    "some-server",
		ResourceGroup: "some-resource-group",
		Subscription:  "some-subscription",
	}

	token := "SomeToken"

	postgres := &Postgres{}
	postgres.SetDBProvider(&mockProvider)
	postgres.SetNext(nil)

	mysql := &Mysql{}
	mysql.SetDBProvider(&mockProvider)
	mysql.SetNext(postgres)

	//When
	err := mysql.AddIp(&serverInfo, token)

	//Then
	assert.Equal(t, nil, err)
}

func TestShouldApplyMysqlProviderWhenIdIsEmptyForRetroCompatibility(t *testing.T) {
	//Given
	mockProvider := mock.MockProvider{}

	mockProvider.On("AddAgentIp").Return(nil)
	mockProvider.On("DeleteAgentIp").Return(nil)

	serverInfo := entity.ServerFirewallIpRule{
		ServerID:      "",
		IP:            "192.168.0.1",
		ServerName:    "some-server",
		ResourceGroup: "some-resource-group",
		Subscription:  "some-subscription",
	}

	token := "SomeToken"

	postgres := &Postgres{}
	postgres.SetDBProvider(&mockProvider)
	postgres.SetNext(nil)

	mysql := &Mysql{}
	mysql.SetDBProvider(&mockProvider)
	mysql.SetNext(postgres)

	//When
	err := mysql.AddIp(&serverInfo, token)

	//Then
	assert.Equal(t, nil, err)
}

func TestShouldReturnErrorWhenThereIsNoMatchID(t *testing.T) {
	//Given
	mockProvider := mock.MockProvider{}

	mockProvider.On("AddAgentIp").Return(nil)
	mockProvider.On("DeleteAgentIp").Return(nil)

	serverInfo := entity.ServerFirewallIpRule{
		ServerID:      "test",
		IP:            "192.168.0.1",
		ServerName:    "some-server",
		ResourceGroup: "some-resource-group",
		Subscription:  "some-subscription",
	}

	token := "SomeToken"

	noProvider := &NoProvider{}

	postgres := &Postgres{}
	postgres.SetDBProvider(&mockProvider)
	postgres.SetNext(noProvider)

	mysql := &Mysql{}
	mysql.SetDBProvider(&mockProvider)
	mysql.SetNext(postgres)

	//When
	err := mysql.AddIp(&serverInfo, token)

	//Then
	assert.Equal(t, "No Provider Match for server some-server server_id test!", err.Error())
}
