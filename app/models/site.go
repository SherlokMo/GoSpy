package models

import (
	"encoding/json"
	"gospy/infrastructure"
	"net/http"
)

type Site struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Interval int    `json:"interval"`
}

func (s Site) ParseJson(w http.ResponseWriter, request *http.Request) Site {
	var site Site
	err := json.NewDecoder(request.Body).Decode(&site)

	infrastructure.ControllerErrorResponder(err, w, http.StatusBadGateway)

	return site

}
