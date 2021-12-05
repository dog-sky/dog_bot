.PHONY: generate
generate:
	protoc --go_out=./pkg/dog --go_opt=paths=source_relative --go-grpc_out=./pkg/dog --go-grpc_opt=paths=source_relative ./api/dog.proto

.PHONY: test
test:
	go test -v -race -timeout 30s -cover ./...

.PHONY: run
run:
	go run cmd/dog/main.go
