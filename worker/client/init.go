package workerClient

import (
	"api-gateway/config"

	"github.com/hibiken/asynq"
)

var client *asynq.Client

func InitializeClient() {
	client = asynq.NewClient(asynq.RedisClientOpt{Addr: config.GetInstance().RedisAddress})
}
