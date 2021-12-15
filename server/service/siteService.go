package service

import (
	"gospy/infrastructure"
	"gospy/models"
)

type SiteService struct {
}

func NewSiteService() *SiteService {
	return &SiteService{}
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
