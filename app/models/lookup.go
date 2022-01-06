package models

import "time"

type LookupResponse struct {
	DNSLookUp      int64     `json:"dnslookuptime"`
	ConnectionTime int64     `json:"connectiontime"`
	TLSHandshake   int64     `json:"tlshandshake"`
	Warning        string    `json:"warning"`
	Status         int       `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type AverageLookup struct {
	DNSLookUp      float32 `json:"avg_dnslookup_time"`
	ConnectionTime float32 `json:"avg_connection_time"`
	TLSHandshake   float32 `json:"avg_tlshandshake_time"`
}
