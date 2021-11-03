## create generated files
go/generate: clean
	@protoc --proto_path=. \
		--go_out=.  \
		--go-grpc_out=require_unimplemented_servers=false:. \
		--go-grpc_opt=paths=source_relative  \
		--go_opt=paths=source_relative \
		grpc.proto

## run the go tests
go/test:
	@-go test ./... --cover -count=1

## lint the go codebase
go/lint:
	@golangci-lint run --config=.golangci.yaml ./...

## run the go server locally as a binary
go/server:
	@go run cmd/server/main.go

.PHONY: test lint generate
