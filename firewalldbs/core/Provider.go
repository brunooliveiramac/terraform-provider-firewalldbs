package core

import "terraform-provider-firewalldbs/firewalldbs/core/entity"

type Provider interface {
	Login(credential *entity.Credential) (token string, err error)
	GetAgentIp() (ip string, err error)
}
