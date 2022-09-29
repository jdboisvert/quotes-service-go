package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	a "github.com/jdboisvert/quotes-service-go/app"
	"github.com/jdboisvert/quotes-service-go/models"
)

var app a.App

func TestMain(m *testing.M) {
	// This is run before the other tests once.
	app = a.App{}

	app.Initialize()

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func quotesEqual(a []models.Quote, b []models.Quote) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !cmp.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestHealth(t *testing.T) {

	req, _ := http.NewRequest("GET", "/health", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var body map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &body)

	if body["status"] != "healthy" {
		t.Errorf("Expected a status of healthy but got %v.", body["status"])
	}
}

func TestGetAllQuotes(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/v1/quotes", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var quotes []models.Quote
	json.Unmarshal(response.Body.Bytes(), &quotes)

	quotesLength := len(quotes)

	if quotesLength != 2 {
		t.Errorf("Expected 2 quotes to be returned but got %v.", quotesLength)
	}

	expectedQuotes := []models.Quote{{Id: "1", Quote: "The greatest glory in living lies not in never falling, but in rising every time we fall.", AuthorName: "Nelson Mandela"}, {Id: "2", Quote: "Life is what happens when you're busy making other plans.", AuthorName: "John Lennon"}}

	if !quotesEqual(quotes, expectedQuotes) {
		t.Errorf("Expected %v to be returned but got %v.", expectedQuotes, quotes)
	}

}

func TestGetQuote(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/v1/quote/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var quote models.Quote
	json.Unmarshal(response.Body.Bytes(), &quote)

	expectedQuote := models.Quote{Id: "1", Quote: "The greatest glory in living lies not in never falling, but in rising every time we fall.", AuthorName: "Nelson Mandela"}

	if !cmp.Equal(expectedQuote, quote) {
		t.Errorf("Expected %v to be returned but got %v.", expectedQuote, quote)
	}

}

func TestGetQuoteNotFound(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/v1/quote/1996", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var body map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &body)

	if body["error"] != "Not Found" {
		t.Errorf("Expected a error message of Not Found but got %v.", body["error"])
	}

}

func TestCreateQuote(t *testing.T) {
	// For simplicity this does check if is in the in memory array.
	payload := []byte(`{"id":"3","quote":"This is a test.", "author_name": "Tester"}`)

	req, _ := http.NewRequest("POST", "/api/v1/quote", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var quote models.Quote
	json.Unmarshal(response.Body.Bytes(), &quote)

	expectedQuote := models.Quote{Id: "3", Quote: "This is a test.", AuthorName: "Tester"}

	if !cmp.Equal(expectedQuote, quote) {
		t.Errorf("Expected to be %v returned but got %v.", expectedQuote, quote)
	}

}

func TestDeleteQuote(t *testing.T) {
	// For simplicity this does check if is removed from the in memory array.
	req, _ := http.NewRequest("DELETE", "/api/v1/quote/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNoContent, response.Code)

	if response.Body.String() != "{}" {
		t.Errorf("Expected to be {} returned but got %v.", response.Body.String())
	}

}

func TestDeleteQuoteNotFound(t *testing.T) {
	// For simplicity this does check if is removed from the in memory array.
	req, _ := http.NewRequest("DELETE", "/api/v1/quote/1996", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var body map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &body)

	if body["error"] != "Not Found" {
		t.Errorf("Expected a error message of Not Found but got %v.", body["error"])
	}

}
