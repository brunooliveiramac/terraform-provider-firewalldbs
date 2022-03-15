package mock

import (
	"github.com/stretchr/testify/mock"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
)

type MockAzureMysqlProvider struct {
	mock.Mock
}

func (mock *MockAzureMysqlProvider) AddAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {
	return nil
}

func (mock *MockAzureMysqlProvider) DeleteAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {
	return nil
}
