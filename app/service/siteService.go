package service

import (
	"gospy/infrastructure"
	"gospy/models"
)

type SiteService struct {
	subscribers map[Observer]bool
	site        *models.Site
}

func NewSiteService() *SiteService {
	siteservice := &SiteService{subscribers: make(map[Observer]bool)}
	siteservice.register(&ScheduleSiteMonitoringService{})
	return siteservice
}
func (p *SiteService) register(o Observer) {
	p.subscribers[o] = true
}

func (p *SiteService) deregister(o Observer) {
	delete(p.subscribers, o)
}

func (p *SiteService) notifyAll() {
	for o, k := range p.subscribers {
		_ = k
		o.Update(map[string]interface{}{
			"id":       p.site.ID,
			"title":    p.site.Title,
			"url":      p.site.Url,
			"interval": p.site.Interval,
		})
	}
}

func (p *SiteService) EmbedSite(s *models.Site) *SiteService {
	p.site = s

	return p
}

func (p SiteService) GetInfo(id interface{}) *models.Site {
	result := infrastructure.Pgsql.QueryRow("SELECT title, url, interval FROM sites WHERE id=$1", id)
	var site models.Site
	err := result.Scan(&site.Title, &site.Url, &site.Interval)
	infrastructure.CheckError(err)
	return &site
}

func (p SiteService) FetchAll() *[]models.Site {
	result := infrastructure.Pgsql.Query("SELECT * FROM sites")
	var sites []models.Site
	for result.Next() {
		var site models.Site
		err := result.Scan(&site.ID, &site.Title, &site.Url, &site.Interval)
		infrastructure.CheckError(err)
		sites = append(sites, site)
	}
	return &sites
}

func (p SiteService) Save() {
	id := infrastructure.Pgsql.Insert("INSERT INTO sites (title, url, interval) VALUES($1, $2, $3)", p.site.Title, p.site.Url, p.site.Interval)
	p.site.ID = id
	p.notifyAll()
}

func (p SiteService) Exists(id interface{}) error {
	return infrastructure.Pgsql.QueryExistance("SELECT id FROM sites WHERE id=$1", id)
}
