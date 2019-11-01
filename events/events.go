package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
type allEvents []event

// To be replace by DB
var dbevents = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
	{
		ID:          "2",
		Title:       "Introduction to Python",
		Description: "Come join us for a chance to learn how Python works and get to eventually try it out",
	},
}

func createNewEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqData, &newEvent)
	dbevents = append(dbevents, newEvent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func getEventbyID(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range dbevents {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range dbevents {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			dbevents = append(dbevents[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range dbevents {
		if singleEvent.ID == eventID {
			dbevents = append(dbevents[:i], dbevents[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dbevents)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// func main() {
// 	fmt.Println("Hello World!")
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", homeLink)
// 	router.HandleFunc("/events", getAllEvents).Methods("GET")
// 	router.HandleFunc("/event", createNewEvent).Methods("POST")
// 	router.HandleFunc("/events/{id}", getEventbyID).Methods("GET")
// 	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
// 	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }
