package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// main to create the api with gorilla mux router and start the server

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/accounts", getAccounts).Methods("GET")
	router.HandleFunc("/api/v1/account/{id}", getAccount).Methods("GET")
	router.HandleFunc("/api/v1/account", createAccount).Methods("POST")
	router.HandleFunc("/api/v1/account/{id}/balance", createAccountBalance).Methods("POST")
	router.HandleFunc("/api/v1/account/{id}/balance", getAccountBalance).Methods("GET")
	router.HandleFunc("/api/v1/account/{id}", deleteAccount).Methods("DELETE")

	// start the server and listen in port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
