# environment variables
dns ?= postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# migration path
migrationPath ?= ./internal/database/migrations

# loading .env
include .env
export $(shell sed 's/=.*//' .env)

# targets
.PHONY: run build seed db-status db-up db-down run-dev clean

run: build
	@./.bin/main
build:
	@go build -o .bin/main cmd/main.go

seed:
	@go run cmd/seed.go

db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(dns) goose -dir=$(migrationPath) status

db-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(dns) goose -dir=$(migrationPath) up

db-down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(dns) goose -dir=$(migrationPath) down

run-dev:
	@air

clean:
	@rm -rf bin

# helper targets
.PHONY: help
help:
	@echo "make run - run the application"
	@echo "make build - build the application"
	@echo "make seed - seed the database"
	@echo "make db-status - check the database status"
	@echo "make db-up - migrate the database up"
	@echo "make db-down - migrate the database down"
	@echo "make run-dev - run the application in development mode"
	@echo "make clean - clean the application"