package repository

import (
	"context"

	"github.com/Khalilkma/search-engine-golang/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pageRepository struct {
	db *pgxpool.Pool
}

func NewPageRepository(db *pgxpool.Pool) PageRepository {
	return &pageRepository{db: db}
}

func (r *pageRepository) Save(ctx context.Context, page *model.Page) error {
	query := `
		INSERT INTO pages (url, title, content)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query, page.URL, page.Title, page.Content)
	return err
}

func (r *pageRepository) Search(ctx context.Context, q string) ([]model.Page, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, url, title, content
		 FROM pages
		 WHERE content ILIKE '%' || $1 || '%'`,
		q,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []model.Page

	for rows.Next() {
		var p model.Page
		err := rows.Scan(&p.ID, &p.URL, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}

	return pages, nil
}
