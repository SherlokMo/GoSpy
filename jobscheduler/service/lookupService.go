package service

import (
	"jobscheduler/helpers"
	"jobscheduler/infrastructure"
	"jobscheduler/models"
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

func (p LookUpService) Save() {
	id := infrastructure.Pgsql.Insert("INSERT INTO lookup (site_id, dnslookup, connection, tlshandshake, warning, status_code) VALUES($1, $2, $3, $4, $5, $6)", p.lookup.Site_id, p.lookup.DNSLookUp, p.lookup.ConnectionTime, p.lookup.TLSHandshake, helpers.NewNullString(p.lookup.Warning), p.lookup.Status)
	p.lookup.ID = id
}
