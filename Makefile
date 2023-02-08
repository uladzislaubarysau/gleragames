.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## run app
	@go run ./cmd/app/main.go

run_dc: ## run app with docker compose
	@docker-compose up --build

run_db: ## run container with db
	@docker-compose up --build db

create_migration: ## add new migration, have to pass parameter "name" 
	@migrate create -ext sql -dir migration -seq $(name)