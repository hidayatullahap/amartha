package database

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	ddl "amartha/db"
	"amartha/generated/sqlc"
)

func Run() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, ddl.DDL); err != nil {
		return err
	}

	queries := sqlc.New(db)

	authors, err := queries.ListLoans(ctx)
	if err != nil {
		return err
	}

	log.Println(authors)
	return nil
}
