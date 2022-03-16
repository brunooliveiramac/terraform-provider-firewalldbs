package mock

import (
	"github.com/stretchr/testify/mock"
	"terraform-provider-firewalldbs/firewalldbs/core/entity"
)

type MockProvider struct {
	mock.Mock
}

func (mock *MockProvider) AddAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {
	return nil
}

func (mock *MockProvider) DeleteAgentIp(firewall *entity.ServerFirewallIpRule, token string) (err error) {
	return nil
}
