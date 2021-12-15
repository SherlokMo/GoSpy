package models

type Site struct {
	ID       int64  `json:"id"`
	Title    string `json:"name"`
	Url      string `json:"url"`
	Interval string `json:"interval"`
}
