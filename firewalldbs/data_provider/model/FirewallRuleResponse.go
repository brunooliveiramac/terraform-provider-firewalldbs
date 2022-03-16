package model

type FirewallRuleResponse struct {
	Properties Properties `json:"properties"`
	Name       string     `json:"name"`
}
