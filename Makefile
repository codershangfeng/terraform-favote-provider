.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: fmt
	mkdir -p ./bin
	go build \
	-o ./bin/terraform-vote-provider ./main.go
	@echo "\033[0;32mSuccessfully build application in ./bin/terraform-vote-provider\033[0m"

.PHONY: run
run: fmt
	go run ./main.go

.PHONY: test
test: fmt 
	go test ./...
