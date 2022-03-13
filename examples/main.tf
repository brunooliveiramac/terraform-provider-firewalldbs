terraform {
  required_providers {
    firewalldbs = {
      version = "1.0.0"
      source  = "hashicorp.com/firewalldbs"
    }
  }
}

provider "firewalldbs" {}

resource "firewalldbs_open" "default" {
  server_name         = "brunoxy-ix4-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxy"
}

resource "firewalldbs_close" "default" {
  server_name         = "brunoxy-ix4-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxy"
  depends_on = [
    firewalldbs_open.default
  ]
}

