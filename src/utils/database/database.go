package database

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	ddl "amartha/db"
	"amartha/generated/sqlc"
	"amartha/src/utils/config"
)

func NewDBConnection(config *config.Config) *sql.DB {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", config.DB.Name)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if _, err := db.ExecContext(ctx, ddl.DDL); err != nil {
		log.Fatalf("Error executing DDL: %v", err)
	}

	return db
}

func NewQueries(db *sql.DB) *sqlc.Queries {
	return sqlc.New(db)
}
