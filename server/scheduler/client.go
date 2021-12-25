package scheduler

import (
	"log"
	"sync"

	"github.com/hibiken/asynq"
)

var locker = &sync.Mutex{}

var Worker *asynq.Client

var redisConnection = asynq.RedisClientOpt{
	Addr: "localhost:6379",
}

func HandleWorker() {
	if Worker == nil {
		locker.Lock()
		defer locker.Unlock()
		Worker = asynq.NewClient(redisConnection)
		log.Println("Worker-client has connected")
	}
}
