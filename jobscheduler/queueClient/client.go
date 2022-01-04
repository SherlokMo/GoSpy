package queueClient

import (
	"log"
	"os"
	"sync"

	"github.com/hibiken/asynq"
)

var locker = &sync.Mutex{}

var Client *asynq.Client

var redisConnection = asynq.RedisClientOpt{
	Addr: os.Getenv("REDIS_ADDR"),
}

func HandleClient() {
	if Client == nil {
		locker.Lock()
		defer locker.Unlock()
		Client = asynq.NewClient(redisConnection)
		log.Println("Worker-client has connected")
	}
}
