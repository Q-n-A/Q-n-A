.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Download and install go mod dependencies
	@go mod download

.PHONY: lint
lint: ## Lint go files
	@golangci-lint run

.PHONY: build
build: ## Build binary
	@go build -o Q-n-A ./*.go

.PHONY: run
run: ## Run Q'n'A directly
	@go run ./*.go

.PHONY: up
up: ## Build and start Q'n'A hot reload environment
	@cd dev && COMPOSE_PROJECT_NAME=q-n-a_hot_reload docker-compose up -d --build

.PHONY: down
down: ## Stop and remove hot reload environment
	@cd dev && docker-compose down

.PHONY: reset-frontend
reset-frontend: stop-front rm-front delete-front-image re-clone-frontend ## Delete frontend container and image completely

.PHONY: re-clone-frontend
re-clone-frontend:
	@cd dev/frontend && rm -rf Q-n-A_UI && git clone https://github.com/Q-n-A/Q-n-A_UI.git && sudo rm -rf dev/frontend/Q-n-A_UI/.git

.PHONY: stop-front
stop-front:
	@docker ps -a | grep Q-n-A_frontend | awk '{print $$1}' | xargs docker stop

.PHONY: rm-front
rm-front:
	@docker ps -a | grep Q-n-A_frontend | awk '{print $$1}' | xargs docker rm

.PHONY: delete-front-image
delete-front-image:
	@docker images -a | grep q-n-a | grep frontend | awk '{print $$3}' | xargs docker rmi

.PHONY: prune
prune: ## Delete redundant images and volumes
	@docker image prune -a && docker volume prune

.PHONY: tbls
tbls: ## Generate tbls DB docs
	@rm -rf docs/db_schema
	@cd docs && tbls doc

.PHONY: grpc
grpc: grpc-go grpc-doc ## Generate gRPC auto-gen files (go & docs)

.PHONY: grpc-go
grpc-go:
	@rm -rf router/grpc
	@mkdir -p router/grpc
	@protoc -I . --go_out=router --go_opt=paths=source_relative --go-grpc_out=router --go-grpc_opt=paths=source_relative **/*.proto

.PHONY: grpc-doc
grpc-doc:
	@protoc -I . --doc_out=docs/grpc.tmpl,grpc.md:docs **/*.proto

.PHONY: chown
chown:
	$(eval name := $(shell whoami))
	@sudo chown -R $(name):$(name) .
