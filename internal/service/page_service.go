package service

import (
	"context"

	"github.com/Khalilkma/search-engine-golang/internal/crawler"
	"github.com/Khalilkma/search-engine-golang/internal/indexer"
	"github.com/Khalilkma/search-engine-golang/internal/model"
	"github.com/Khalilkma/search-engine-golang/internal/repository"
)

type PageService struct {
	Repo  repository.PageRepository
	Index *indexer.InvertedIndex
}

func NewPageService(r repository.PageRepository) *PageService {

	idx := indexer.New()

	return &PageService{
		Repo:  r,
		Index: idx,
	}
}

func (s *PageService) Search(ctx context.Context, query string) ([]*model.Page, error) {

	ids := s.Index.Search(query)

	// simplified version, still search in the database
	results, err := s.Repo.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	_ = ids // gonna use later

	return results, nil
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

		s.Index.Add(page)
	}

	return pages, nil
}
