package service

import (
	"database/sql"
	"gospy/infrastructure"
	"gospy/models"
)

type LookUpService struct{}

func NewLookUpService() *LookUpService {
	return &LookUpService{}
}

func (p LookUpService) FetchScopeSite(site_id string) *[]models.LookupResponse {
	result := infrastructure.Pgsql.Query("SELECT dnslookup, connection, tlshandshake, warning, status_code, created_at FROM lookup WHERE site_id=$1 ORDER BY id DESC LIMIT 50", site_id)
	var lookups []models.LookupResponse
	for result.Next() {
		var lookup models.LookupResponse
		var warning sql.NullString
		err := result.Scan(&lookup.DNSLookUp, &lookup.ConnectionTime, &lookup.TLSHandshake, &warning, &lookup.Status, &lookup.CreatedAt)
		infrastructure.CheckError(err)
		if warning.Valid {
			lookup.Warning = warning.String
		}
		lookups = append(lookups, lookup)
	}
	return &lookups

}

func (p LookUpService) GetAverageScopeSite(site_id string) *models.AverageLookup {
	result := infrastructure.Pgsql.QueryRow("SELECT avg(l.dnslookup) AS avg_dnslookup, avg(l.connection) AS avg_connection, avg(l.tlshandshake) AS avg_tlshandshake FROM (SELECT dnslookup, connection, tlshandshake FROM lookup WHERE site_id=$1 AND warning IS NULL ORDER BY id DESC LIMIT 50) AS l", site_id)
	var average models.AverageLookup
	err := result.Scan(&average.DNSLookUp, &average.ConnectionTime, &average.TLSHandshake)
	infrastructure.CheckError(err)

	return &average
}
