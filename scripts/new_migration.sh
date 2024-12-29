#!/bin/bash

# Exit on error 
set -e

# Exit if no file variable provided.
if [ -z "$1" ]; then 
    echo "[Error] No name provided for migration file."
    echo "[Solution] make new_migration file=<name>"
    exit 1
fi

# Migrations directory.
MIGRATIONS_DIR="./internal/db/migrations"

goose -dir "$MIGRATIONS_DIR" create -s "$1" sql 

exit 0