package models

type Site struct {
	ID       int64  `json:"id"`
	Title    string `json:"name"`
	Email    string `json:"email"`
	Url      string `json:"url"`
	Interval string `json:"interval"`
}
