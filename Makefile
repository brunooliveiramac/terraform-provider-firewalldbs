TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=edu
NAME=azurepostgressql
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=linux_amd64

build:
	go get -d -v ./... && go mod vendor && go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

apply:
	rm -rf examples/.terraform && rm -rf examples/.terraform.lock.hcl && rm -rf examples/.terraform.lock.hcl \
 	&& rm -rf examples/terraform.tfstate && cd examples && terraform init && terraform plan && terraform apply -auto-approve \

local: install apply