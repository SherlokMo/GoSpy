package main

import (
	"jobscheduler/infrastructure"
	"jobscheduler/queueClient"
	"jobscheduler/scheduler"
)

func main() {
	infrastructure.HandlePostgre()
	queueClient.HandleClient()
	scheduler.RunScheduler()
}
