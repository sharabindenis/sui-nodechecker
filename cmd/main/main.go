package main

import (
	"github.com/gorilla/mux"
	"github.com/sharabindenis/sui-nodechecker/pkg/controllers"
	"github.com/sharabindenis/sui-nodechecker/pkg/routes"
	"log"
	"net/http"
)

func main() {
	go controllers.TelegramBot()
	r := mux.NewRouter()
	routes.RegisterNodeRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
