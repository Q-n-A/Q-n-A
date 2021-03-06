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

.PHONY: serve
serve: ## Run Q'n'A server directly
	@go run ./*.go serve

.PHONY: healthcheck
healthcheck: ## Healthcheck Q'n'A server
	@go run ./*.go healthcheck

.PHONY: config
config: ## Display current server configuration
	@go run ./*.go config

.PHONY: up
up: ## Build and start Q'n'A hot reload environment
	@cd dev && COMPOSE_PROJECT_NAME=q-n-a_hot_reload docker-compose up -d --build

.PHONY: down
down: ## Stop and remove hot reload environment
	@cd dev && docker-compose down

.PHONY: reset-frontend
reset-frontend: stop-front rm-front delete-front-image re-clone-frontend ## Delete frontend image and re-clone frontend repo to update frontend container

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

.PHONY: gen
gen: ## Generate go auto-gen files
	@GOFLAGS=-mod=mod go generate ./...

.PHONY: grpc
grpc: grpc-go grpc-doc ## Generate gRPC auto-gen files

.PHONY: grpc-go
grpc-go:
	@rm -rf server/protobuf
	@mkdir -p server/protobuf
	@protoc -I . --go_out=server --go_opt=paths=source_relative --go-grpc_out=server --go-grpc_opt=paths=source_relative protobuf/*.proto

.PHONY: grpc-doc
grpc-doc:
	@protoc -I . --doc_out=docs/settings/grpc.tmpl,grpc.md:docs/api protobuf/*.proto

.PHONY: grpc-list
grpc-list: ## List up gRPC services
	@grpcurl -plaintext :9001 list

.PHONY: grpc-ping
grpc-ping: ## Ping to gRPC server
	@grpcurl -plaintext :9001 grpc.Ping/Ping

.PHONY: chown
chown:
	$(eval name := $(shell whoami))
	@sudo chown -R $(name):$(name) .
