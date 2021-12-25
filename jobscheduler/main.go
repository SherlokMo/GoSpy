package main

import (
	"jobscheduler/infrastructure"
	"jobscheduler/scheduler"
)

func main() {
	infrastructure.HandlePostgre()
	scheduler.RunScheduler()
}
