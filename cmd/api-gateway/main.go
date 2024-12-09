package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"api-gateway/config"
	workerClient "api-gateway/worker/client"
	workerServer "api-gateway/worker/server"
)

func main() {
	config.Init()
	serverConfig := config.GetInstance()

	// middlewares.InitSentry(serverConfig.SentryEnabled)

	if len(os.Args) > 1 && os.Args[1] == "worker" {
		startWorker(serverConfig)
	} else {
		startApp(serverConfig)
	}
}

func startApp(serverConfig *config.Config) {
	workerClient.InitializeClient()
	app := NewApp(serverConfig) // Ensure this function is correctly defined in app.go
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	if err := app.Start(ctx); err != nil {
		fmt.Println("failed to start app:", err)
	}
}

func startWorker(serverConfig *config.Config) {
	workerServer := workerServer.NewWorkerServer(serverConfig)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	if err := workerServer.Start(ctx); err != nil {
		fmt.Println("failed to start worker server:", err)
	}
}
