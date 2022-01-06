package middlewares

import (
	"gospy/infrastructure"
	"gospy/service"
	"net/http"

	"github.com/gorilla/mux"
)

func SiteExists(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		siteService := service.NewSiteService()
		if err := siteService.Exists(vars["site"]); err != nil {
			infrastructure.JsonResponse(map[string]interface{}{
				"message":              "No such site with this id",
				"additional_resources": "[GET]/sites",
			}, w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
