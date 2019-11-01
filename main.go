package main

import (
	"log"
	"net/http"
)

func main() {
	r := routes.getRouter()
	// fmt.Println("Hello World!")
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homeLink)
	// router.HandleFunc("/events", getAllEvents).Methods("GET")
	// router.HandleFunc("/event", createNewEvent).Methods("POST")
	// router.HandleFunc("/events/{id}", getEventbyID).Methods("GET")
	// router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	// router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
