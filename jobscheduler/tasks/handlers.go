package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"jobscheduler/infrastructure"
	"jobscheduler/models"
	"jobscheduler/service"
	"log"
	"net/http"

	"github.com/hibiken/asynq"
)

// Each handler will run on it's own worker
func HandleWebsiteMonitor(c context.Context, t *asynq.Task) error {
	var site models.Site
	err := json.Unmarshal(t.Payload(), &site)
	fmt.Println(site)
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
		LookUpService.EmbedModel(&lookup).Save()
		return nil
	}
	req, err := http.Head(site.Url)
	_ = err
	lookup.Status = req.StatusCode
	LookUpService.EmbedModel(&lookup).Save()
	log.Println("Saved", site.Url, "With lookup Id", lookup.ID)
	return nil
}