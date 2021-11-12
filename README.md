# Simple gRPC set up in Go

A simple Go application to use for reference, incorporating most of the basic principles of application development.

## Implemented Technologies/Concepts
* HTTP server
* gRPC: client and server in Go using generated files
* Docker
  * Multi-stage builds
  * Docker Compose
    * external dependencies
      * Database
      * Vault
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
* [Go playground](https://play.golang.org/) - beware time always starts at 2009-11-10 23:00:00 UTC!
* [Go by example](https://gobyexample.com/) - good 'How to' guides for common patterns
* [Concurrency Talk by Rob Pike](https://talks.golang.org/2012/concurrency.slide#1) - Rob Pikes talks on understand concurrency 

## Generating Go code from the proto file
Go and the [protobuf](https://google.golang.org/protobuf) package allow us to define our Protobuf message and services in 
.proto files, as it is the standard, and then generate the Go code and interfaces required to implement it.
I am deliberately ignoring the generated files in git so the user can ensure their environment can correctly generate them.  
To generate the files you can run the following command, assuming you have installed the dependencies detailed below.  

```protoc --proto_path=. --go_out=. --go-grpc_out=require_unimplemented_servers=false:.  --go-grpc_out=. --go-grpc_opt=paths=source_relative  --go_opt=paths=source_relative grpc.proto```  

Alternatively this command has been added to `make generate` 

## Principles of software development
* [S.O.L.I.D.](https://en.wikipedia.org/wiki/SOLID) principles of software development
* [12 Factor Apps](https://12factor.net/)
 
## Tools
[gRPC UI](https://github.com/fullstorydev/grpcui) - handy tool for locally testing gRPC services.
[Postman](https://www.postman.com/downloads/) - handy tool for locally testing HTTP services, import the collection and environment from the root folder.
[JWT.io](https://jwt.io/) - for exploring JWT tokens, token never leaves the browser so safe to use.  

### Requirements

I've tried to keep the requirements as a check in the Makefile, run `make requirements` to check the required applications are available.

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


# Todos
* Run generate in docker image 
* Add helm charts to deploy to a minikube instance
* Use information in vault to get the jwt token
* CI/CD tooling for GitLab/Github - I'm hosting in github, but we are targeting gitlab 
* Update HTTP server to provide same functionality as the gRPC service
* Add gRPC interceptors to match http middleware
* Redis?
* Logging with Zap
* Viper
* Cobra?