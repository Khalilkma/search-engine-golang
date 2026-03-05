package crawler

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Khalilkma/search-engine-golang/internal/model"
	"github.com/PuerkitoBio/goquery"
)

func Crawler(seed string, maxDepth int) ([]*model.Page, error) {
	visited := make(map[string]bool)
	seed = normalizeURL(seed)
	queue := []string{seed}
	var results []*model.Page

	parsedSeed, err := url.Parse(seed)
	if err != nil {
		return nil, err
	}

	baseDomain := parsedSeed.Host

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// BFS
	for range maxDepth {

		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {

			if len(queue) == 0 {
				break
			}

			currentURL := normalizeURL(queue[0])
			queue = queue[1:]

			if visited[currentURL] {
				continue
			}
			visited[currentURL] = true

			resp, err := client.Get(currentURL)
			if err != nil {
				continue
			}

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			resp.Body.Close()
			if err != nil {
				continue
			}

			// Extract data
			title := strings.TrimSpace(doc.Find("title").Text())

			var description string
			if desc, exists := doc.Find("meta[name='description']").Attr("content"); exists {
				description = strings.TrimSpace(desc)
			}

			var headings []string
			doc.Find("h1").Each(func(i int, s *goquery.Selection) {
				text := strings.TrimSpace(s.Text())
				if text != "" {
					headings = append(headings, text)
				}
			})

			content := strings.TrimSpace(doc.Find("body").Text())

			results = append(results, &model.Page{
				URL:         currentURL,
				Title:       title,
				Description: description,
				Headings:    strings.Join(headings, ", "),
				Content:     content,
			})

			// Extract links
			doc.Find("a").Each(func(i int, s *goquery.Selection) {
				href, exists := s.Attr("href")
				if !exists {
					return
				}

				link, err := url.Parse(href)
				if err != nil {
					return
				}

				base, _ := url.Parse(currentURL)
				resolved := base.ResolveReference(link)

				if resolved.Host == baseDomain {
					normalized := normalizeURL(resolved.String())
					queue = append(queue, normalized)
				}
			})
		}
	}

	return results, nil
}

func normalizeURL(raw string) string {

	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}

	// remove fragment (#section)
	u.Fragment = ""

	// remove final /
	if u.Path == "/" {
		u.Path = ""
	}

	return u.String()
}
