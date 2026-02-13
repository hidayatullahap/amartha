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

	db, err := sql.Open("sqlite", config.DB.Name)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if _, err := db.ExecContext(ctx, ddl.DDL); err != nil {
		log.Fatalf("Error executing DDL: %v", err)
	}

	q := sqlc.New(db)
	hasSeed, _ := q.GetFlag(ctx, HAS_SEED)
	if hasSeed.Value != "true" {
		if _, err := db.ExecContext(ctx, ddl.Seeds); err != nil {
			log.Fatalf("Error executing seed: %v", err)
		}

		q.InsertFlag(ctx, sqlc.InsertFlagParams{Key: HAS_SEED, Value: "true"})
	}

	return db
}

func NewQueries(db *sql.DB) sqlc.Querier {
	return sqlc.New(db)
}
