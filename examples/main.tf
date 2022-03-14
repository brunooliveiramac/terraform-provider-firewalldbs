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
  server_name         = "brunoxyz-k8j-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxyz"
}

resource "firewalldbs_close" "default" {
  server_name         = "brunoxyz-k8j-north-eu-sandbox"
  resource_group_name = "bees-eu-sbx-brunoxyz"

  depends_on = [
    firewalldbs_open.default
  ]
}

