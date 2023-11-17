package asynqclient

import "github.com/hibiken/asynq"

const redisAddr = "127.0.0.1:6379"

func NewClient() *asynq.Client {
	return asynq.NewClient(
		asynq.RedisClientOpt{Addr: redisAddr},
	)
}
