default: help

-include make/Makefile.*

## check requirements are installed, offers link to install page if missing
requirements:
	@$(MAKE) -s -f make/Makefile.golang _go/requirements
	@$(MAKE) -s -f make/Makefile.docker _docker/requirements
	@$(MAKE) -s -f make/Makefile.vault _vault/requirements
	@$(MAKE) -s -f make/Makefile.database _database/requirements
	@$(MAKE) -s -f make/Makefile.k8s _k8s/requirements

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