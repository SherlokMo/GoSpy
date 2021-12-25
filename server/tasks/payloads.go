package tasks

import (
	"encoding/json"
	"gospy/infrastructure"
	"gospy/models"

	"github.com/hibiken/asynq"
)

const (
	TypeMonitorWebsite = "website:monitor"
)

func NewWebsiteMonitor(site models.Site) *asynq.Task {
	bytePayload, err := json.Marshal(site)
	infrastructure.CheckError(err)
	return asynq.NewTask(TypeMonitorWebsite, bytePayload)
}
