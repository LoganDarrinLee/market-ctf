// For running database migrations
package main

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/LoganDarrinLee/market-ctf/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	// All migration logs will go to logs/migartions.log
	log.Println("Starting database migration.")

	// New environment configuration
	env := config.NewEnv()
	os.Setenv("PGPASSWORD", env.DBPassword)

	// New context
	ctx := context.Background()

	// Database connection
	dbpool := config.InitDB(ctx, env)
	defer dbpool.Close()

	// Migrate database
	migrateDB(dbpool)

	// Export schema for sqlc
	schemafile := "./internal/db/schema/schema.sql"
	if err := exportSchema(env, schemafile); err != nil {
		log.Fatal(err)
	}

	// Successful database migration.
	log.Println("Database successfully migrated.")
}

// Migrate with goose. Takes connection pool and converts into *database/sql.DB
func migrateDB(dbpool *pgxpool.Pool) {
	// Set goose dialect to postgres
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	// Convert database connection pool to *sql.DB for goose
	db := stdlib.OpenDBFromPool(dbpool)
	if err := goose.Up(db, "./internal/db/migrations"); err != nil {
		log.Fatal(err)
	}

	// Explicitly close db connection, not the pool
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

// Export our new schema file for sqlc to run.
func exportSchema(env *config.EnvironmentConfig, schemafile string) error {
	// Define pg_dump command with args
	cmd := exec.Command(
		"pg_dump",
		"-U", env.DBUser,
		"-h", env.DBHost,
		"-p", env.DBPort,
		"-s", env.DBName,
	)

	// Setup the file to dump the schema into
	outputFile, err := os.Create(schemafile)
	if err != nil {
		log.Println("Erorr with schema file: ", err)
		return err
	}
	defer outputFile.Close()
	cmd.Stdout = outputFile

	// Execute command
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
