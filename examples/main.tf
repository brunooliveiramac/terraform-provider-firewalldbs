terraform {
  required_providers {
    azurepostgressql = {
      version = "0.2"
      source  = "hashicorp.com/edu/azurepostgressql"
    }
  }
}

provider "azurepostgressql" {
  host     = "bees-eastus2-pim-sandbox.postgres.database.azure.com"
  port     = 5432
  database = "pim"
  username = ""
  password = ""
}

resource "azurepostgressql_role" "default" {
  name        = "manager1"
  database    = "pim"
  schema      = "public"
  tables      = ["ALL"]
  privileges  = ["SELECT", "DELETE"]
}

resource "azurepostgressql_user" "brunom" {
  username     = "pirulito2"
  password     = "kYsun3sAUG456Xu9"
  role         = azurepostgressql_role.default.name
}

