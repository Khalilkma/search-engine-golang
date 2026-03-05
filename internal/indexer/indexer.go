package indexer

import (
	"strings"

	"github.com/Khalilkma/search-engine-golang/internal/model"
)

/*
This file implements a simple inverted index for the search engine.

The idea is simple: when a page is added, the text is analyzed and split into tokens.
Each token is mapped to the page ID where it appears.
So when searching, the query is analyzed the same way and the matching pages are returned.
*/

type InvertedIndex struct {
	index map[string][]int
}

func New() *InvertedIndex {
	return &InvertedIndex{
		index: make(map[string][]int),
	}
}

func (idx *InvertedIndex) Add(page *model.Page) {

	doc := buildDoc(page)

	tokens := Analyze(doc)

	seen := make(map[string]bool)

	for _, token := range tokens {

		if seen[token] {
			continue
		}

		idx.index[token] = append(idx.index[token], page.ID)

		seen[token] = true
	}
}

func buildDoc(p *model.Page) string {

	var builder strings.Builder

	builder.WriteString(p.Title)
	builder.WriteString(" ")
	builder.WriteString(p.Description)
	builder.WriteString(" ")
	builder.WriteString(p.Headings)
	builder.WriteString(" ")
	builder.WriteString(p.Content)

	return builder.String()
}

func (idx *InvertedIndex) Search(query string) []int {

	tokens := Analyze(query)

	result := map[int]int{}

	for _, token := range tokens {

		docs := idx.index[token]

		for _, id := range docs {
			result[id]++
		}
	}

	var ranked []int

	for id := range result {
		ranked = append(ranked, id)
	}

	return ranked
}
