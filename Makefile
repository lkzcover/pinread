LOCAL_BIN := $(CURDIR)/bin
LOCALDB_COMPOSE := docker-compose --project-name pinread --file $(CURDIR)/docker-compose.yml
MIGRATION_DIR := $(CURDIR)/migrations

LOCAL_DB_NAME := pinread
LOCAL_DB_PORT := 5432
LOCAL_DB_CONNECTION := postgresql://postgres:postgres@localhost:$(LOCAL_DB_PORT)/$(LOCAL_DB_NAME)?sslmode=disable

bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.11.2
	GOBIN=$(LOCAL_BIN) go install github.com/go-jet/jet/v2/cmd/jet@latest

db-create:
	psql -p $(LOCAL_DB_PORT) -U postgres -c "drop database if exists $(LOCAL_DB_NAME)"
	psql -p $(LOCAL_DB_PORT) -U postgres -c "create database $(LOCAL_DB_NAME)"

db-reset: db-create db-up

# Применение всех доступных и еще не принятых миграций
db-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} -allow-missing postgres $(LOCAL_DB_CONNECTION) up

# Local environment start&stop
local-env-start:
	$(LOCALDB_COMPOSE) up -d
local-env-stop:
	$(LOCALDB_COMPOSE) down


jet:
	$(LOCAL_BIN)/jet -dsn=$(LOCAL_DB_CONNECTION) -schema=public -path=./internal/generated

local-run:
	cd ./cmd && export $(shell cat ./deploy/local.env | xargs) && go run ./main.go