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
  server_id           = "/subscriptions/5c92b4a1-d813-42e0-804d-0c0e64218b27/resourceGroups/bees-eu-sbx-brunoxyy/providers/Microsoft.DBforMySQL/servers/brunoxyy-6fu-north-eu-sandbox"
  server_name         = "brunoxyy-6fu-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxyy"
}

resource "firewalldbs_close" "default" {
  server_id           = "/subscriptions/5c92b4a1-d813-42e0-804d-0c0e64218b27/resourceGroups/bees-eu-sbx-brunoxyy/providers/Microsoft.DBforMySQL/servers/brunoxyy-6fu-north-eu-sandbox"
  server_name         = "brunoxyy-6fu-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxyy"

  depends_on = [
    firewalldbs_open.default
  ]

}

