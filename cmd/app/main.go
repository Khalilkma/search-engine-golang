package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Khalilkma/search-engine-golang/database"
	"github.com/Khalilkma/search-engine-golang/internal/handler"
	"github.com/Khalilkma/search-engine-golang/internal/repository"
	"github.com/Khalilkma/search-engine-golang/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(" .env não encontrado")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	defer db.Close()

	log.Println("Conectado ao banco")

	// Injeção de dependência
	pageRepo := repository.NewPageRepository(db)
	pageService := service.NewPageService(pageRepo)
	pageHandler := handler.NewPageHandler(pageService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Rotas
	r.GET("/search", pageHandler.Search)

	// Rotas crawl
	r.POST("/crawl", pageHandler.Crawl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Servidor rodando na porta", port)
	r.Run(":" + port)
}
