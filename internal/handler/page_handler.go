package handler

import (
	"github.com/Khalilkma/search-engine-golang/internal/service"
	"github.com/gin-gonic/gin"
)

type PageHandler struct {
	Service *service.PageService
}

func NewPageHandler(svc *service.PageService) *PageHandler {
	return &PageHandler{Service: svc}
}

func (h *PageHandler) Crawl(c *gin.Context) {

	type CrawlRequest struct {
		URL   string `json:"url"`
		Depth int    `json:"depth"`
	}

	var req CrawlRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	if req.URL == "" {
		c.JSON(400, gin.H{"error": "url is required"})
		return
	}

	if req.Depth <= 0 {
		req.Depth = 1
	}

	pages, err := h.Service.CrawlAndSave(c.Request.Context(), req.URL, req.Depth)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, pages)
}

func (h *PageHandler) Search(c *gin.Context) {

	query := c.Query("q")
	if query == "" {
		c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
		return
	}

	results, err := h.Service.Search(c.Request.Context(), query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, results)
}
