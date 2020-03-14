package http

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	Routes "github.com/jdschrack/mongotutorial/http/routes"
)

var Router mux.Router
var port string = ":9001"

func ConfigureHttp() {
	log.Println("Configuring HTTP Server")
	var router *mux.Router = mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(log.Writer(), router)
	router.HandleFunc("/person", Routes.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/person", Routes.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", Routes.GetPersonEndpoint).Methods("GET")
	http.ListenAndServe(port, loggedRouter)
	log.Println("HTTP Host listening on port " + port)
}