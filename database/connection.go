package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the pq driver
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

func Connect() error {
	dsn := "postgres://postgres:password@localhost:5432/project?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	pgconn := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	sqldb := sql.OpenDB(pgconn)

	db = bun.NewDB(sqldb, pgdialect.New())

	// Enable query debugging
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return nil
}

func Initialize() error {
	dsn := "postgres://postgres:password@postgresql-postgresql-ha-pgpool.loopabord.svc.cluster.local:5432/postgres?sslmode=disable"
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres database: %w", err)
	}
	defer sqldb.Close()

	// Check if the project database exists
	var exists bool
	err = sqldb.QueryRow("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = 'project')").Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if project database exists: %w", err)
	}

	// Create the project database if it doesn't exist
	if !exists {
		_, err = sqldb.Exec("CREATE DATABASE project")
		if err != nil {
			return fmt.Errorf("failed to create project database: %w", err)
		}
		log.Println("Database project created successfully.")
	}

	return nil
}

// Ensure you close the db connection when the application shuts down
func Close() error {
	return db.Close()
}
