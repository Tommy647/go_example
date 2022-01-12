# Build the binary
FROM golang:1.17-alpine AS builder

WORKDIR /app

# copy just the mod and sum file, these do not change often, but require
# the most work to process (downloads) so we can cache this layer separately
COPY go.mod go.sum ./
# get the modules
RUN go mod download

# copy the remaining files
COPY . .
ADD /db/networkDBMigrations /db/networkDBMigrations
# disable CGO, forces go to build staticly without clibs
ENV CGO_ENABLED 0


# simple ci/cd check
RUN go test ./...
# build the jwt server binary
RUN go build -o jwt cmd/jwtProvider/main.go
# build the http server binary
RUN go build -o httpserver cmd/httpserver/main.go
# build the grpcServer binary
RUN go build -o grpcserver cmd/grpcServer/main.go
# build the networkMGMT binary
RUN go build -o networkMGMT cmd/networkMGMT/main.go

# create a fresh image, without the go toolset
FROM alpine:3.14 AS grpcserver
# copy over the binary we built above
COPY --from=builder /app/grpcserver .
# expose our working port
EXPOSE 9090
# command we run when starting the container
CMD ./grpcserver

# create a fresh image, without the go toolset
FROM alpine:3.14 AS httpserver
# copy over the binary we built above
COPY --from=builder /app/httpserver .
# expose our working port
EXPOSE 8080
# command we run when starting the container
CMD ./httpserver

# create a fresh image, without the go toolset
FROM alpine:3.14 as jwtserver
# copy over the binary we built above
COPY --from=builder /app/jwt .
# expose our working port
EXPOSE 8081
# command we run when starting the container
CMD ./jwt

# create a fresh image, without the go toolset
FROM alpine:3.14 AS networkMGMT
# copy over the binary we built above
COPY --from=builder /app/networkMGMT .
COPY --from=builder /db/networkDBMigrations /db/networkDBMigrations
# expose our working port
EXPOSE 9091
# command we run when starting the container
CMD ./networkMGMT