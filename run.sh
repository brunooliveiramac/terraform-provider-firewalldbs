go get -d -v ./... \
  && go mod vendor \
  && go build -o terraform-provider-azurepostgressql  \
  && mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/azurepostgressql/0.2/linux_amd64 \
  && mv terraform-provider-azurepostgressql ~/.terraform.d/plugins/hashicorp.com/edu/azurepostgressql/0.2/linux_amd64 \
  && rm -rf examples/.terraform \
  && rm -rf examples/.terraform.lock.hcl \
  && cd examples \
  && terraform init && terraform plan