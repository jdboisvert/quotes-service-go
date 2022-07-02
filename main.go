package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getQuotes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all quotes")
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting a quote")
}

func createQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a quote")
}

func updateQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upating a quote")
}

func deleteQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting a quote")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/quote", getQuotes).Methods("GET")
	r.HandleFunc("/api/v1/quote/{id}", getQuote).Methods("GET")
	r.HandleFunc("/api/v1/quote", createQuote).Methods("POST")
	r.HandleFunc("/api/v1/quote/{id}", updateQuote).Methods("PUT")
	r.HandleFunc("/api/v1/quote/{id}", deleteQuote).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
