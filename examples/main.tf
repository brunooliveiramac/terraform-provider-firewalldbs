terraform {
  required_providers {
    firewalldbs = {
      version = "1.0.1"
      source  = "hashicorp.com/brunooliveiramac/firewalldbs"
    }
  }
}

provider "firewalldbs" {

}

resource "firewalldbs_open" "default" {
  server_name         = "brunoxya-rs6-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxya"
}

resource "firewalldbs_close" "default" {
  server_name         = "brunoxya-rs6-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxya"

  depends_on = [
    firewalldbs_open.default
  ]
}

