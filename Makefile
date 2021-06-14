HOSTNAME=registry.terraform.io
NAME=favote
NAMESPACE=codershangfeng
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
LOCAL_MIRROR_DIR=${HOME}/.terraform.d/plugins

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: fmt
	mkdir -p ./bin
	go build \
	-o ./bin/${BINARY} ./main.go
	@echo "\033[0;32mSuccessfully build application in ./bin/${BINARY}\033[0m"

.PHONY: run
run: fmt
	go run ./main.go

.PHONY: test
test: fmt 
	go test ./...

.PHONY: install
install: fmt build
	mkdir -p ${LOCAL_MIRROR_DIR}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/darwin_amd64/
	@mv ./bin/${BINARY} ${LOCAL_MIRROR_DIR}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/darwin_amd64/${BINARY}_${VERSION}
	@echo "\033[0;32mSuccessfully install local provider\033[0m"