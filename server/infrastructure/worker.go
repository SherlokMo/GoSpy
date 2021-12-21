package infrastructure

import (
	"log"
	"sync"

	"github.com/hibiken/asynq"
)

var locker = &sync.Mutex{}

var worker *asynq.Client

var redisConnection = asynq.RedisClientOpt{
	Addr: "redis:6379",
}

func HandleWorker() {
	if worker == nil {
		locker.Lock()
		defer locker.Unlock()
		worker = asynq.NewClient(redisConnection)
		log.Println("Worker-client has connected")
	}
}
