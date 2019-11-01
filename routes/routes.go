package routes

import (
	"github.com/gorilla/mux"

	"github.com/rvats/go-rest-api/events"
)

func getRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", events.homeLink)
	// router.HandleFunc("/events", events.getAllEvents).Methods("GET")
	// router.HandleFunc("/event", events.createNewEvent).Methods("POST")
	// router.HandleFunc("/events/{id}", events.getEventbyID).Methods("GET")
	// router.HandleFunc("/events/{id}", events.updateEvent).Methods("PATCH")
	// router.HandleFunc("/events/{id}", events.deleteEvent).Methods("DELETE")

	return router

	// log.Fatal(http.ListenAndServe(":8080", router))
}
