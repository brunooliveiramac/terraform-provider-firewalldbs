terraform {
  required_providers {
    firewalldbs = {
      version = "1.0.1"
      source  = "hashicorp.com/brunooliveiramac/firewalldbs"
    }
  }
}

provider "firewalldbs" {}

resource "firewalldbs_open" "default" {
  server_id           = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/batata-sbx-brunoxyy/providers/Microsoft.DBforMySQL/servers/batata-6fu-sbx"
  server_name         = "batata-6fu-sbx"
  resource_group_name = "batata-sbx-brunoxyy"
}

resource "firewalldbs_close" "default" {
  server_id           = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/batata-sbx-brunoxyy/providers/Microsoft.DBforMySQL/servers/batata-6fu-sbx"
  server_name         = "batata-6fu-sbx"
  resource_group_name = "batata-sbx-brunoxyy"

  depends_on = [
    firewalldbs_open.default
  ]
}

