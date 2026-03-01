package repository

import (
	"context"

	"github.com/Khalilkma/search-engine-golang/internal/model"
)

type PageRepository interface {
	Save(ctx context.Context, page *model.Page) error
	Search(ctx context.Context, query string) ([]model.Page, error)
}
