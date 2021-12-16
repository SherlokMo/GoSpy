package handlers

import (
	"gospy/infrastructure"
	"gospy/service"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleSitesRequest(r *mux.Router) {
	r.HandleFunc("/sites", getSites)
}

func getSites(w http.ResponseWriter, r *http.Request) {
	SiteService := service.NewSiteService()
	sites := SiteService.FetchAll()

	infrastructure.JsonResponse(sites, w)
}
