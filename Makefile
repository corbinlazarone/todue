.PHONY: help test dev dev-build dev-stop dev-restart dev-logs build run db-up db-down db-reset db-shell migrate-up migrate-down migrate-status clean

help: 
	@echo "todue - Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  make dev            - Start all services (web-client, server, db)"
	@echo "  make dev-build      - Rebuild images and start services (use after code changes)"
	@echo "  make dev-stop       - Stop all services"
	@echo "  make dev-restart    - Restart containers (no rebuild)"
	@echo "  make dev-logs       - View logs from all services"
	@echo "  make dev-logs-server - View server logs only"
	@echo ""
	@echo "Local Development (outside Docker):"
	@echo "  make run            - Run server locally (requires db running)"
	@echo "  make build          - Build the server binary (runs tests first)"
	@echo "  make test           - Run server tests"
	@echo ""
	@echo "Database:"
	@echo "  make db-up          - Start only the database"
	@echo "  make db-down        - Stop the database"
	@echo "  make db-reset       - Reset database (destroy and recreate)"
	@echo "  make db-shell       - Open psql shell in database"
	@echo ""
	@echo "Migrations:"
	@echo "  make migrate-up     - Run all pending migrations"
	@echo "  make migrate-down   - Rollback last migration"
	@echo "  make migrate-status - Show migration status"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean          - Remove built binaries and docker volumes"

# Testing
test:
	@echo "Running server tests..."
	cd server && go test ./...

# Development - Docker Compose
dev:
	@echo "Starting all services..."
	docker-compose up -d
	@echo "Services started! Web: http://localhost:3000, API: http://localhost:4000"

dev-build:
	@echo "Rebuilding images and starting services..."
	docker-compose up -d --build
	@echo "Services rebuilt and started! Web: http://localhost:3000, API: http://localhost:4000"

dev-stop:
	@echo "Stopping all services..."
	docker-compose down

dev-restart:
	@echo "Restarting all services..."
	docker-compose restart

dev-logs:
	docker-compose logs -f

dev-logs-server:
	docker-compose logs -f server

# Local Development (run server outside Docker)
run:
	@echo "Running server locally..."
	@if [ ! -f .env ]; then echo "Error: .env file not found"; exit 1; fi
	@echo "Loading environment from .env (replacing 'db' with 'localhost')..."
	@export $$(cat .env | grep -v '^#' | xargs) && \
	export DB_CONN=$$(echo $$DB_CONN | sed 's/@db:/@localhost:/') && \
	cd server && go run ./cmd/api

build:
	@echo "Running tests..."
	cd server && go test ./...
	@echo "Building server binary..."
	cd server && go build -o ../bin/todue-server ./cmd/api
	@echo "Binary created at: bin/todue-server"

# Database Operations
db-up:
	@echo "Starting database only..."
	docker-compose up -d db
	@echo "Waiting for database to be ready..."
	@sleep 2
	@echo "Database ready at: localhost:5432"

db-down:
	@echo "Stopping database..."
	docker-compose stop db

db-reset:
	@echo "⚠️  WARNING: This will destroy all data!"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		echo "Resetting database..."; \
		docker-compose down -v; \
		docker-compose up -d db; \
		echo "Waiting for database to be ready..."; \
		sleep 3; \
		echo "Database reset complete!"; \
	else \
		echo "Cancelled."; \
	fi

db-shell:
	@echo "Opening psql shell..."
	@docker-compose exec db psql -U $$(grep POSTGRES_USER .env | cut -d '=' -f2) -d $$(grep POSTGRES_DB .env | cut -d '=' -f2)

# Migration Operations (using goose)
migrate-up:
	@echo "Running migrations..."
	cd server/cmd/internals/migrations && goose postgres "$$(grep DB_CONN ../../../../.env | cut -d '=' -f2- | sed 's/@db:/@localhost:/')" up

migrate-down:
	@echo "Rolling back last migration..."
	cd server/cmd/internals/migrations && goose postgres "$$(grep DB_CONN ../../../../.env | cut -d '=' -f2- | sed 's/@db:/@localhost:/')" down

migrate-status:
	@echo "Migration status:"
	cd server/cmd/internals/migrations && goose postgres "$$(grep DB_CONN ../../../../.env | cut -d '=' -f2- | sed 's/@db:/@localhost:/')" status

# Cleanup
clean:
	@echo "Cleaning up..."
	rm -rf bin/
	docker-compose down -v
	@echo "Cleanup complete!"
