package models

type Site struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Interval int    `json:"interval"`
}
