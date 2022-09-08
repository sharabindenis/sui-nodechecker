package routes

import (
	"github.com/gorilla/mux"
	"github.com/sharabindenis/sui-nodechecker/pkg/controllers"
)

var RegisterNodeRoutes = func(router *mux.Router) {
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./pkg/views/")))
	//router.HandleFunc("/", controllers.Index)
	//router.HandleFunc("/node/", controllers.GetNodes).Methods("GET")
	router.HandleFunc("/start/", controllers.CreateSchedule).Methods("POST")
	router.HandleFunc("/stop/", controllers.Stop).Methods("GET")
	router.HandleFunc("/jobs/", controllers.ShowJobs).Methods("GET")
	//router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	//router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
