.PHONY: setup generate run build clean

APP_NAME = main
BINARY = ./bin/$(APP_NAME)
MAIN = cmd/server/main.go

generate:
	@echo ">> Generating Swagger docs..."
	@swag init -d $(shell find . -name '*.go' -exec dirname {} \; | sort -u | tr '\n' ',' | sed 's/,$$//')

setup:
	@echo ">> Setting up environment..."
	@chmod +x ./scripts/*
	@./scripts/setup.sh

build:
	@echo ">> Building binary..."
	@go build -o $(BINARY) $(MAIN)

run: build
	@echo ">> Running application..."
	@$(BINARY)

clean:
	@echo ">> Cleaning up..."
	@rm -rf $(BINARY) ./docs ./bin
