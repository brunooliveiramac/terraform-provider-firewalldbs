terraform {
  required_providers {
    firewalldbs = {
      version = "1.0.0"
      source  = "hashicorp.com/firewalldbs"
    }
  }
}

provider "firewalldbs" {
    client_id         = ""
    client_secret     = ""
    subscription_id   = ""
    tenant_id         = ""
    agent_ip          = ""
}

resource "firewalldbs_open" "default" {
  server_name         = "brunoxy-ix4-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxy"
  agent_ip            = "192.168.1.1"
}

resource "firewalldbs_close" "default" {
  server_name         = "brunoxy-ix4-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxy"
  agent_ip            = "192.168.1.1"

  depends_on = [
    firewalldbs_open.default
  ]
}

