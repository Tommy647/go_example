default: help

-include make/Makefile.*

## check requirements are installed, offers link to install page if missing
requirements:
	@-which go || echo go missing
	@-which golangci-lint || echo golangci-lint missing: https://golangci-lint.run/usage/install/
	@-which protoc || echo protoc missing: https://grpc.io/docs/protoc-installation/
	@-which protoc-gen-go || echo protoc-gen-go missing:  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@-which protoc-gen-go-grpc || echo protoc-gen-go-grpc missing:  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	@-which docker || echo docker missing:
	@-docker compose stop || echo docker compose missing: https://github.com/docker/compose/tree/v2
	@-which vault || echo valut missing: https://learn.hashicorp.com/tutorials/vault/getting-started-install
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