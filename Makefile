SHELL := /bin/bash

setup: 
	source development.env && docker compose up -d db 	# Start the Postgres container

new_migration:
	scripts/new_migration.sh $(file) 

migrate:
	scripts/migrate.sh

run: 
	source development.env && go run cmd/server/main.go 	# Run the app

stop:
	docker compose down 	# Stop and remove the container