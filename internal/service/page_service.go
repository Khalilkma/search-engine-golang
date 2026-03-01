package service

import (
	"context"

	"github.com/Khalilkma/search-engine-golang/internal/model"
	"github.com/Khalilkma/search-engine-golang/internal/repository"
)

type PageService struct {
	repo repository.PageRepository
}

func NewPageService(repo repository.PageRepository) *PageService {
	return &PageService{repo: repo}
}

func (s *PageService) Create(ctx context.Context, page *model.Page) error {
	return s.repo.Save(ctx, page)
}

func (s *PageService) Search(ctx context.Context, query string) ([]model.Page, error) {
	return s.repo.Search(ctx, query)
}
