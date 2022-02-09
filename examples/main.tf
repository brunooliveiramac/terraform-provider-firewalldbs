terraform {
  required_providers {
    azurepostgressql = {
      version = "0.2"
      source  = "hashicorp.com/edu/azurepostgressql"
    }
  }
}

provider "azurepostgressql" {
  host            = "postgres_server_ip"
  port            = 5432
  database        = "postgres"
  username        = "postgres_user"
  password        = "postgres_password"
}

resource "azurepostgressql_user" "edu" {
  items {
    coffee {
      id = 3
    }
    quantity = 2
  }
  items {
    coffee {
      id = 2
    }
    quantity = 2
  }
}
