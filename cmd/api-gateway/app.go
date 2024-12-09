package main

import (
	"context"
	// "fmt"
	"net/http"
	"time"

	"api-gateway/config"
	"api-gateway/middlewares"
	"api-gateway/routes"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type App struct {
	config *config.Config
	router http.Handler
	rdb    *redis.Client
	logger *zap.Logger
}

func NewApp(cfg *config.Config) *App {
	logger, _ := zap.NewProduction()
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddress,
	})

	router := routes.NewRoutes()
	router = middlewares.RequestLogger(logger)(router)

	return &App{
		config: cfg,
		router: router,
		rdb:    rdb,
		logger: logger,
	}
}

func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: app.router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.Fatal("Server failed", zap.Error(err))
		}
	}()

	<-ctx.Done()
	app.logger.Info("Shutting down server...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctxShutdown)
}
