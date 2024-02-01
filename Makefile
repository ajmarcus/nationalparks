.PHONY: help run

help:
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)


run: ## Run fetch.go
	source .env && go run fetch.go
