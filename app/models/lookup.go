package models

import "time"

type Lookup struct {
	ID             int64     `json:"id"`
	Site_id        int64     `json:"site_id"`
	DNSLookUp      int64     `json:"dnslookuptime"`
	ConnectionTime int64     `json:"connectiontime"`
	TLSHandshake   int64     `json:"tlshandshake"`
	Warning        string    `json:"warning",omitempty`
	Status         int       `json:"status",omitempty`
	CreatedAt      time.Time `json:"created_at"`
}

type AverageLookup struct {
	DNSLookUp      int64 `json:"avg_dnslookup_time"`
	ConnectionTime int64 `json:"avg_connection_time"`
	TLSHandshake   int64 `json:"avg_tlshandshake_time"`
}
