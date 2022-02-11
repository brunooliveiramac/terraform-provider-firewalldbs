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
  database = "postgres"
  username = ""
  password = ""
}

resource "azurepostgressql_role" "default" {
  name        = "read_write"
  privileges  = ["SELECT", "DELETE"]
}

resource "azurepostgressql_user" "brunom" {
  username     = "brunom"
  password     = "brunom"
  database     = "pim"
  role          = azurepostgressql_role.default.name
}

resource "azurepostgressql_user" "bastiao" {
  username     = "bastiao"
  password     = "bastiao"
  database     = "postgres"
  role         = azurepostgressql_role.default.name
}
