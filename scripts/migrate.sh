#!/bin/bash

# Exit on error 
set -e

# Set our development environment up
source development.env 

# Run our migration script, log output to termianl & migrations log with tee
go run cmd/migrations/migrate.go > >(tee -a ./logs/migrations.log) 2>&1

# If migration unsuccessful, exit to avoid errors with sqlc generation.
if [ $? -eq 0 ]; then
    echo "Succesful database migration"
else 
    echo "[Error] Failed to migrate database."
    exit 1
fi 

sqlc generate