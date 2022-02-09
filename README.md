
Terraform Provider for Azure Porsgres Sql Database Operations

$ go mod vendor

$ go build -o terraform-provider-azure-postgres-user

$ export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"

$ mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/azure-postgres-user/0.2/$OS_ARCH

$ mv terraform-provider-azure-postgres-user ~/.terraform.d/plugins/hashicorp.com/edu/azure-postgres-user/0.2/$OS_ARCH

$ terraform init && terraform apply --auto-approve


