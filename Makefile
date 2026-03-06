# Environment variables — no .env file needed
export PORT       := 8080
export APP_ENV    := local
export DB_HOST    := localhost
export DB_PORT    := 27017
export DB_USERNAME := appuser
export DB_PASSWORD := apppassword
export DB_NAME    := mydb

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Run migrations
migrate-up:
	@echo "Running migrations..."
	@go run cmd/migrate/main.go up

migrate-down:
	@echo "Rolling back migrations..."
	@go run cmd/migrate/main.go down

# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
		air; \
		echo "Watching...";\
	else \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air; \
			echo "Watching...";\
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi

format:
	@echo "Formatting code..."
	@go fmt ./...

.PHONY: all build run test clean watch docker-run docker-down itest migrate-up migrate-down