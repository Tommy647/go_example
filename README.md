# Simple gRPC set up in Go

A simple Go application to use for reference, incorporating most of the basic principles of application development.

## Implemented Technologies/Concepts
* HTTP server
* gRPC: client and server in Go using generated files
* Docker
  * Multi-stage builds
  * Docker Compose
    * TODO: add external dependencies
    * TODO: add the client app
* Go
  * Project layout
    * Domain based
    * Multi Entry point Monolith
  * Basic concurrency
  * Interfaces
  * Middleware (http and gRPC)
  * TODO: Logger - production level:
  * Database

## Makefile
All the commands needed to run the application are documented in the Makefile and its children, run `make help` for details  

## Guides for Go
* [playground](https://play.golang.org/) - beware time always starts at 2009-11-10 23:00:00 UTC!
* [Go by example](https://gobyexample.com/)
* [Concurrency Talk by Rob Pike](https://talks.golang.org/2012/concurrency.slide#1)

## Generating Go code from the proto file
Go and the [protobuf](https://google.golang.org/protobuf) package allow us to define our Protobuf message and services in 
.proto files, as it is the standard, and then generate the Go code and interfaces required to implement it.
I am deliberately ignoring the generated files in git so the user can ensure their environment can correctly generate them.  
To generate the files you can run the following command, assuming you have installed the dependencies detailed below.  
`protoc --proto_path=. --go_out=.  --go-grpc_out=. --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative grpc.proto`b  
Alternatively this command has been added to `make generate` 
## Principles of software development
* [S.O.L.I.D.](https://en.wikipedia.org/wiki/SOLID) principles of software development
* [12 Factor Apps](https://12factor.net/)
 
## Tools
[gRPC UI](https://github.com/fullstorydev/grpcui) - handy tool for locally testing a gRPC service

### Requirements
* [Go](https://golang.org/) install and go/bin in $PATH 
* Go Plugins
  * [golangci-lint](https://golangci-lint.run/) `go install go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
  * [gRPC](https://pkg.go.dev/google.golang.org/grpc) `go install google.golang.org/grpc@latest`
  * [protoc-gen-go](https://developers.google.com/protocol-buffers/docs/overview) `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  * [protoc-gen-go-grpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc) `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`
* [protoc](https://grpc.io/docs/protoc-installation/)
* [Docker](https://www.docker.com/)
  * [Docker Compose Plugin](https://github.com/docker/compose/tree/v2)
* TODO:  [minikube](https://minikube.sigs.k8s.io/docs/start/)
* [Vault](https://learn.hashicorp.com/tutorials/vault/getting-started-install)
* [DB Migrate](https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md) `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1`
