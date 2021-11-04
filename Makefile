default: help

-include make/Makefile.*

## check requirements are installed, offers link to install page if missing
requirements:
	@-which go &>1 || echo golang
	@-which protoc &>1 || echo https://grpc.io/docs/protoc-installation/
	@-which protoc-gen-go &>1 || echo go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@-which protoc-gen-go-grpc &>1 || echo go google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@-which docker &>1

## run all targets, as a quick smoke test
all: clean go/generate go/lint go/test docker/start docker/stop

## This help screen
help:
	@printf "Available targets:\n\n"
	@-awk '/^[a-zA-Z\-\\_0-9%:\\]+/ { \
	  helpMessage = match(lastLine, /^## (.*)/); \
	  if (helpMessage) { \
		helpCommand = $$1; \
		helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
	gsub("\\\\", "", helpCommand); \
	gsub(":+$$", "", helpCommand); \
		printf "  \x1b[32;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
	  } \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"

.PHONY: help