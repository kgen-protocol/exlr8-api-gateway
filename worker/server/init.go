package workerServer

import (
	"api-gateway/config"
	"context"

	"github.com/hibiken/asynq"
)

type WorkerServer struct {
	server *asynq.Server
}

func NewWorkerServer(cfg *config.Config) *WorkerServer {
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: cfg.RedisAddress},
		asynq.Config{
			Concurrency: 10, // Adjust based on your app's needs
		},
	)

	return &WorkerServer{server: server}
}

func (ws *WorkerServer) Start(ctx context.Context) error {
	mux := asynq.NewServeMux()
	// Example task handler
	// mux.HandleFunc("task_name", YourTaskHandlerFunction)

	return ws.server.Run(mux)
}
