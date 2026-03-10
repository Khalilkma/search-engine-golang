package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(pool *pgxpool.Pool) error {

	query := `
	CREATE TABLE IF NOT EXISTS pages (
		id SERIAL PRIMARY KEY,
		url TEXT UNIQUE,
		title TEXT,
		description TEXT,
		headings TEXT,
		content TEXT
	);`

	_, err := pool.Exec(context.Background(), query)

	return err
}