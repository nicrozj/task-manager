package config

import (
	"embed"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrationFS embed.FS

func NewDBClient() (*sqlx.DB, error) {
	fmt.Println(Envs.DB_URL)
	db, err := sqlx.Connect("postgres", Envs.DB_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := MigrateDB(db); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return db, nil
}

func MigrateDB(db *sqlx.DB) error {
	log.Println("Launching migrations...")

	sqlBytes, err := migrationFS.ReadFile("migrations/schema.sql")
	if err != nil {
		return fmt.Errorf("couldn't read the migration file: %w", err)
	}

	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		return fmt.Errorf("migration execution error: %w", err)
	}

	log.Println("Migrations have been successfully applied")
	return nil
}
