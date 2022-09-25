package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdboisvert/quotes-service-go/models"
	"github.com/jdboisvert/quotes-service-go/utils"
)

var quotes = []models.Quote{
	{Id: "1", Quote: "The greatest glory in living lies not in never falling, but in rising every time we fall.", AuthorName: "Nelson Mandela"},
	{Id: "2", Quote: "Life is what happens when you're busy making other plans.", AuthorName: "John Lennon"},
}

func heathCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Performing health check")
	utils.RespondWithJSON(w, 200, models.Status{Status: "healthy"})
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all quotes")
	utils.RespondWithJSON(w, 200, quotes)
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Getting quote with id", id)

	for _, quote := range quotes {
		// For the sake of simplicity since this is just a list in memory. So O(n) complexity however can be O(1) with a hashmap.
		if quote.Id == id {
			utils.RespondWithJSON(w, 200, quote)
			return
		}
	}

	log.Println("Unable to find quote with id", id)
	utils.RespondWithError(w, 404, "Not Found")

}

func createQuote(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating a quote")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var quote models.Quote
	json.Unmarshal(reqBody, &quote)

	quotes = append(quotes, quote)

	utils.RespondWithJSON(w, 201, quote)
}

func updateQuote(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating a quote")

}

func deleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("Attempting to delete quote with id", id)

	for index, quote := range quotes {
		if quote.Id == id {
			quotes = append(quotes[:index], quotes[index+1:]...)
			utils.RespondWithJSON(w, 204, nil)
		}

	}

	log.Println("Quote does not exist to delete", id)
	utils.RespondWithError(w, 404, "Not Found")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/health", heathCheck).Methods("GET")

	router.HandleFunc("/api/v1/quotes", getQuotes).Methods("GET")

	router.HandleFunc("/api/v1/quote", createQuote).Methods("POST")
	router.HandleFunc("/api/v1/quote/{id}", getQuote).Methods("GET")
	router.HandleFunc("/api/v1/quote/{id}", deleteQuote).Methods("DELETE")

	log.Println("Listening to port 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
