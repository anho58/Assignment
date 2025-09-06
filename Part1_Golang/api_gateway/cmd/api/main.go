package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"part1_golang/api_gateway/internal/application"
	domain "part1_golang/api_gateway/internal/domain/book"
	httpHanlder "part1_golang/api_gateway/internal/infrastructure/http"
	memory "part1_golang/api_gateway/internal/infrastructure/memory"
	_ "part1_golang/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book API
// @version 1.0
// @description This is a sample server for managing books.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	repo := memory.NewBookRepo()
	_, _ = repo.Add(context.Background(), domain.New(uuid.Nil, "Book 1", "Book 1 Author", "Description Test", 50))
	_, _ = repo.Add(context.Background(), domain.New(uuid.Nil, "Book 2", "Book 1 Author", "", 50))

	service := application.NewBookService(repo)

	// transport layer
	handler := httpHanlder.NewBookHandler(service)

	// gin setup
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
		handler.RegisterRoutes(api)
	}

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
