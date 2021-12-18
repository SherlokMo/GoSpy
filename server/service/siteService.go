package service

import (
	"gospy/infrastructure"
	"gospy/models"
)

type SiteService struct {
	site *models.Site
}

func NewSiteService() *SiteService {
	return &SiteService{}
}

func (p *SiteService) EmbedSite(s *models.Site) *SiteService {
	p.site = s

	return p
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
}
