package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	manage "github.com/fiap-postech-soat1-group21/product-api/product-api/adapter/handler"
	"github.com/fiap-postech-soat1-group21/product-api/product-api/adapter/repository"
	"github.com/fiap-postech-soat1-group21/product-api/product-api/internal/domain/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

const (
	shutdownTimeout = 5 * time.Second
	pathPrefix      = "/api/v1"
)

var httpPort = fmt.Sprintf(":%s", os.Getenv("API_PORT"))

func main() {

	repository := repository.NewRepository()
	m := manage.New(
		&manage.UseCases{
			Product: usecase.NewProductUseCase(repository.Product),
		},
	)

	engine := gin.Default()

	v1Routes := engine.Group(pathPrefix)

	m.RegisterRoutes(v1Routes)

	engine.Run(httpPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: engine,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
	}
}
