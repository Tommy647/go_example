default: help

-include make/Makefile.*

## check requirements are installed, offers link to install page if missing
requirements:
	@-which go > /dev/null || echo go missing
	@-which protoc > /dev/null || echo protoc missing: https://grpc.io/docs/protoc-installation/
	@-which protoc-gen-go > /dev/null|| echo protoc-gen-go missing:  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@-which protoc-gen-go-grpc > /dev/null || echo protoc-gen-go-grpc missing:  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@-which docker > /dev/null || echo docker missing:
	@-which vault > /dev/null || echo vault missing: https://www.vaultproject.io/docs/install
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