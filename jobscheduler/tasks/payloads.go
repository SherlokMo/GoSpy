package tasks

import (
	"encoding/json"
	"jobscheduler/infrastructure"

	"github.com/hibiken/asynq"
)

const (
	TypeMonitorWebsite = "website:monitor"
)

func NewWebsiteMonitor(payload map[string]interface{}) *asynq.Task {

	bytePayload, err := json.Marshal(payload)
	infrastructure.CheckError(err)
	return asynq.NewTask(TypeMonitorWebsite, bytePayload)
}
