package main

import (
	handler "gospy/handlers"
	"gospy/infrastructure"
	"gospy/scheduler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(mux.CORSMethodMiddleware(router))

	handler.HandleSitesRequest(router)
	handler.HandleLookupRequest(router)

	log.Println("Listening at ::8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	infrastructure.HandlePostgre()
	scheduler.HandleWorker()
	handleRequests()
}
