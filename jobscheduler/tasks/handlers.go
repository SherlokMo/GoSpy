package tasks

import (
	"context"
	"encoding/json"
	"jobscheduler/infrastructure"
	"jobscheduler/models"
	"jobscheduler/queueClient"
	"jobscheduler/service"
	"log"
	"net/http"
	"time"

	"github.com/hibiken/asynq"
)

// Each handler will run on it's own worker
func HandleWebsiteMonitor(c context.Context, t *asynq.Task) error {
	var site models.Site
	err := json.Unmarshal(t.Payload(), &site)
	infrastructure.CheckError(err)
	tracer, err := infrastructure.TraceHttpConnection("HEAD", site.Url)
	LookUpService := service.NewLookUpService()
	var lookup models.Lookup = models.Lookup{
		Site_id:        site.ID,
		DNSLookUp:      tracer.DNS.Milliseconds(),
		ConnectionTime: tracer.Connection.Milliseconds(),
		TLSHandshake:   tracer.TLSHandshake.Milliseconds(),
	}
	if err != nil {
		lookup.Warning = err.Error()
	}
	req, err := http.Head(site.Url)
	if err == nil {
		lookup.Status = req.StatusCode
	}
	LookUpService.EmbedModel(&lookup).Save()
	log.Println("Scanned ", site.Url, "With lookup Id", lookup.ID)
	queueClient.Client.Enqueue(t, asynq.ProcessIn(time.Duration(site.Interval)*time.Minute))
	return nil
}
