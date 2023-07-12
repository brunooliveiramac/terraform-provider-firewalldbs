package entity

type ServerFirewallIpRule struct {
	ServerID      string `json:"server_id"`
	IP            string `json:"ip"`
	ServerName    string `json:"server"`
	ResourceGroup string `json:"resource_group"`
	Subscription  string `json:"subscription"`
	IsFlexible    bool   `json:"is_flexible"`
}
