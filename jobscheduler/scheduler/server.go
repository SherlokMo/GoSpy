package scheduler

import (
	"jobscheduler/tasks"
	"os"

	"github.com/hibiken/asynq"
)

func RunScheduler() {
	redisConnection := asynq.RedisClientOpt{
		Addr: os.Getenv("REDIS_ADDR"), // Redis server address
	}
	worker := asynq.NewServer(redisConnection, asynq.Config{
		// how many concurrent workers to use.
		Concurrency: 10,
		Queues: map[string]int{
			"default": 10,
		},
	})
	mux := asynq.NewServeMux()
	prepareHandlers(mux)

	if err := worker.Run(mux); err != nil {
		panic(err)
	}
}

func prepareHandlers(mux *asynq.ServeMux) {
	mux.HandleFunc(tasks.TypeMonitorWebsite, tasks.HandleWebsiteMonitor)
}
