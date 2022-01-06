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
	checkRequester.HandleFunc("/checks/{site:[0-9]+}", getByLookup).Methods("GET")
}

func getByLookup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	lookUpService := service.NewLookUpService()

	response := map[string]interface{}{
		"lookups": lookUpService.FetchScopeSite(vars["site"]),
	}

	infrastructure.JsonResponse(response, w)
}
