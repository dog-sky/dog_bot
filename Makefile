.PHONY: generate
generate:
	protoc --go_out=./pkg/dog --go_opt=paths=source_relative --go-grpc_out=./pkg/dog --go-grpc_opt=paths=source_relative ./api/dog.proto

.PHONY: test
test:
	go test -v -race -timeout 30s -cover ./...

.PHONY: run
run:
	go run cmd/dog/main.go

# generates mocks
.PHONY: generate-mocks
generate-mocks:
	find . -name '*_minimock.go' -delete
	go generate ./...
	go mod tidy -compat=1.17
