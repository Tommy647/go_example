
docker:
	docker compose up -d --build

docker-stop:
	docker compose down

docker-logs:
	docker compose logs -f -t

run-server:
	go run cmd/server/main.go

generate: clean
	@protoc --proto_path=. \
		   --go_out=.  \
		   --go-grpc_out=require_unimplemented_servers=false:. \
		   --go-grpc_opt=paths=source_relative  \
		   --go_opt=paths=source_relative \
		   grpc.proto

test:
	@go test ./... --cover -count=1

grpcui:
	grpcui -plaintext localhost:9090



clean:
	@-rm *.pb.go 2> /dev/null ||: