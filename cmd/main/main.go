package main

import (
	"github.com/gorilla/mux"
	"github.com/sharabindenis/sui-nodechecker/pkg/models"
	"github.com/sharabindenis/sui-nodechecker/pkg/routes"
	"log"
	"net/http"
)

var abc map[models.Schedule]string

func main() {
	r := mux.NewRouter()
	routes.RegisterNodeRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
