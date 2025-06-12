.PHONY: run build clean test docker-up docker-down

# Default MySQL credentials - can be overridden with environment variables
DB_HOST ?= localhost
DB_USER ?= root
DB_PASSWORD ?= root
DB_NAME ?= userdb
DB_PORT ?= 3306
PORT ?= 8080

# Build the application
build:
	go build -o main .

# Clean build artifacts
clean:
	rm -f main
	rm -f userapi

# Run tests
test:
	go test ./... -v

# Run the application locally
run: build
	DB_HOST=$(DB_HOST) \
	DB_USER=$(DB_USER) \
	DB_PASSWORD=$(DB_PASSWORD) \
	DB_NAME=$(DB_NAME) \
	DB_PORT=$(DB_PORT) \
	PORT=$(PORT) \
	./main

# Start the application with Docker Compose
docker-up:
	docker-compose up --build

# Stop Docker Compose services
docker-down:
	docker-compose down

# Initialize/Reset the database
init-db:
	mysql -h $(DB_HOST) -u $(DB_USER) -p$(DB_PASSWORD) < schema.sql

# Help command
help:
	@echo "Available commands:"
	@echo "  make build       - Build the application"
	@echo "  make clean       - Remove build artifacts"
	@echo "  make test        - Run tests"
	@echo "  make run         - Run the application locally"
	@echo "  make docker-up   - Start the application with Docker Compose"
	@echo "  make docker-down - Stop Docker Compose services"
	@echo "  make init-db     - Initialize/Reset the database"
	@echo ""
	@echo "Environment variables (current values):"
	@echo "  DB_HOST     = $(DB_HOST)"
	@echo "  DB_USER     = $(DB_USER)"
	@echo "  DB_PASSWORD = $(DB_PASSWORD)"
	@echo "  DB_NAME     = $(DB_NAME)"
	@echo "  DB_PORT     = $(DB_PORT)"
	@echo "  PORT        = $(PORT)" 