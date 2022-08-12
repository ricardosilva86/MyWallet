SHELL := zsh
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

server:
	@echo "Running server..."
	@go run main.go server
.PHONY : server

migrate:
	@echo "Running migrations..."
	@go run main.go migrations
.PHONY : migrate

rename-local:
	@echo "Running local DB server..."
	@cp app.local.yaml app.yaml
.PHONY : rename-local

rename-remote:
	@echo "Running remote DB server..."
	@cp app.remote.yaml app.yaml
.PHONY : rename-remote

local: rename-local server
.PHONY : local

remote: rename-remote server
.PHONY : remote