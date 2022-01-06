package service

import (
	"gospy/infrastructure"
	"gospy/models"
)

type LookUpService struct {
	lookup *models.Lookup
}

func NewLookUpService() *LookUpService {
	return &LookUpService{}
}

func (p *LookUpService) EmbedModel(s *models.Lookup) *LookUpService {
	p.lookup = s
	return p
}

func (p LookUpService) FetchScopeSite(site_id string) *[]models.Lookup {
	result := infrastructure.Pgsql.Query("SELECT dnslookup, connection, tlshandshake, warning, status_code, created_at FROM lookup WHERE site_id=$1", site_id)
	var lookups []models.Lookup
	for result.Next() {
		var lookup models.Lookup
		err := result.Scan(&lookup.DNSLookUp, &lookup.ConnectionTime, &lookup.TLSHandshake, &lookup.Warning, &lookup.Status, lookup.CreatedAt)
		infrastructure.CheckError(err)
		lookups = append(lookups, lookup)
	}
	return &lookups

}
