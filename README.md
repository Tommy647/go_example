# Simple gRPC set up in Go

A simple set up of gRPC in Go.

I am deliberately ignoring the generated files in git so the user can ensure their environment can correctly generate them.

## Generating Go from the proto file
`protoc --proto_path=. --go_out=.  --go-grpc_out=. --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative grpc.proto`

### Tools
[gRPC UI](https://github.com/fullstorydev/grpcui) - handy tool for locally testing a gRPC service

### Requirements
* Go install and go/bin in $PATH 
* Go Plugins 
  * go install google.golang.org/grpc@latest
    * go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
* [protoc](https://grpc.io/docs/protoc-installation/)
* Docker