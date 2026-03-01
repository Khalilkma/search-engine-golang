package handler

import (
	"net/http"

	"github.com/Khalilkma/search-engine-golang/internal/model"
	"github.com/Khalilkma/search-engine-golang/internal/service"
	"github.com/gin-gonic/gin"
)

type PageHandler struct {
	service *service.PageService
}

func NewPageHandler(service *service.PageService) *PageHandler {
	return &PageHandler{service: service}
}

// POST /pages
func (h *PageHandler) Create(c *gin.Context) {
	var page model.Page

	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.Create(c.Request.Context(), &page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "page created successfully",
	})
}

// GET /search?q=algo
func (h *PageHandler) Search(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter 'q' is required",
		})
		return
	}

	results, err := h.service.Search(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}
