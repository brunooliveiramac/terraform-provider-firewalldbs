TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
NAME=firewalldbs
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=linux_amd64

build:
	go get -d -v ./... && go mod vendor && go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAME}/${VERSION}/${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

apply:
	rm -rf examples/.terraform && rm -rf examples/.terraform.lock.hcl && rm -rf examples/.terraform.lock.hcl \
 	&& rm -rf examples/terraform.tfstate && cd examples && terraform init && terraform plan && terraform apply -auto-approve \

local: install apply

fmt:
	gofmt -w $(GOFMT_FILES)