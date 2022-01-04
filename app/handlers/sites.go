package handlers

import (
	"gospy/infrastructure"
	"gospy/models"
	"gospy/service"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleSitesRequest(r *mux.Router) {
	r.HandleFunc("/sites", getSites).Methods("GET")
	r.HandleFunc("/sites", addSite).Methods("POST")
}

func getSites(w http.ResponseWriter, r *http.Request) {
	SiteService := service.NewSiteService()
	sites := SiteService.FetchAll()

	infrastructure.JsonResponse(sites, w)
}

func addSite(w http.ResponseWriter, r *http.Request) {

	site := models.Site{}.ParseJson(w, r)
	SiteService := service.NewSiteService().EmbedSite(&site)
	SiteService.Save()
	response := map[string]interface{}{
		"message": "Your server was added to the list and will be monitored",
		"site":    site,
	}

	infrastructure.JsonResponse(response, w)
}
