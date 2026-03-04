package service

import (
	"context"

	"github.com/Khalilkma/search-engine-golang/internal/crawler"
	"github.com/Khalilkma/search-engine-golang/internal/model"
	"github.com/Khalilkma/search-engine-golang/internal/repository"
)

type PageService struct {
	Repo repository.PageRepository
}

func NewPageService(r repository.PageRepository) *PageService {
	return &PageService{Repo: r}
}

func (s *PageService) Search(ctx context.Context, query string) ([]*model.Page, error) {
	return s.Repo.Search(ctx, query)
}

func (s *PageService) CrawlAndSave(ctx context.Context, url string, depth int) ([]*model.Page, error) {

	pages, err := crawler.Crawler(url, depth)
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		err := s.Repo.Save(ctx, page)
		if err != nil {
			return nil, err
		}
	}

	return pages, nil
}
