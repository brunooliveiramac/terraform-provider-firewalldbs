
# Firewalldbs Provider


The Firewaldbs providers is intended to interact with all database firewall from Azure. This is needed to add
an ip in order to be able to create users and granting permissions using other custom provider.


## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "firewalldbs" {
    client_id         = ""
    client_secret     = ""
    subscription_id   = ""
    tenant_id         = ""
    agent_ip          = "192.168.1.1"
}

resource "firewalldbs_open" "default" {
  server_name         = "batata-6fu-sbx"
  resource_group_name = "batata-sbx-brunoxyy
  agent_ip            = "192.168.1.1"
}

resource "firewalldbs_close" "default" {
  server_name         = "batata-6fu-sbx"
  resource_group_name = "batata-sbx-brunoxyy"
  agent_ip            = "192.168.1.1"

  depends_on = [
    firewalldbs_open.default
  ]
}
```