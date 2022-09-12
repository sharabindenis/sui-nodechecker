package main

import (
	"github.com/sharabindenis/sui-nodechecker/pkg/controllers"
)

func main() {
	controllers.TelegramBot()
	//r := mux.NewRouter()
	//routes.RegisterNodeRoutes(r)
	//http.Handle("/", r)
	//log.Fatal(http.ListenAndServe("localhost:9010", r))
}
