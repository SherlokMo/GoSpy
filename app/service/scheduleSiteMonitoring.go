package service

import (
	"gospy/models"
	"gospy/scheduler"
	"gospy/tasks"
	"log"
	"time"

	"github.com/hibiken/asynq"
	"github.com/mitchellh/mapstructure"
)

type ScheduleSiteMonitoringService struct{}

func (s ScheduleSiteMonitoringService) Update(payload map[string]interface{}) {
	var site models.Site
	mapstructure.Decode(payload, &site)

	taskinfo, err := scheduler.Worker.Enqueue(tasks.NewWebsiteMonitor(site), asynq.ProcessIn(time.Duration(site.Interval)*time.Minute))
	log.Println("Scheduled with id:", taskinfo.ID, " Url: ", site.Url, " with error: ", err)
}
