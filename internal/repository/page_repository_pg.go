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
		INSERT INTO pages (url, title, description, headings, content)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (url) DO NOTHING
	`

	_, err := r.db.Exec(ctx, query,
		page.URL,
		page.Title,
		page.Description,
		page.Headings,
		page.Content,
	)

	return err
}

func (r *pageRepository) Search(ctx context.Context, query string) ([]*model.Page, error) {

	sql := `
		SELECT id, url, title, description, headings, content
		FROM pages
		WHERE title ILIKE '%' || $1 || '%'
		   OR description ILIKE '%' || $1 || '%'
		   OR content ILIKE '%' || $1 || '%'
	`

	rows, err := r.db.Query(ctx, sql, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []*model.Page

	for rows.Next() {
		var p model.Page
		err := rows.Scan(
			&p.ID,
			&p.URL,
			&p.Title,
			&p.Description,
			&p.Headings,
			&p.Content,
		)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &p)
	}

	return pages, nil
}
