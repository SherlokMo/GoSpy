package handlers

import (
	middlewares "gospy/Middlewares"
	"gospy/infrastructure"
	"gospy/service"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleLookupRequest(r *mux.Router) {
	checkRequester := r.Methods(http.MethodGet).Subrouter()
	checkRequester.Use(middlewares.SiteExists)
	checkRequester.HandleFunc("/checks/{site:[0-9]+}", getByLookup)
}

func getByLookup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	lookUpService := service.NewLookUpService()
	sitesService := service.NewSiteService()
	site := sitesService.GetInfo(vars["site"])
	response := map[string]interface{}{
		"site": map[string]interface{}{
			"title": site.Title,
			"url":   site.Url,
		},
		"lookups":      lookUpService.FetchScopeSite(vars["site"]),
		"average_time": lookUpService.GetAverageScopeSite(vars["site"]),
	}

	infrastructure.JsonResponse(response, w)
}
