package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// main to create the api with gorilla mux router and start the server
func main() {

	// create the router
	router := mux.NewRouter()
	// attach the handler to the router
	//Get all accounts
	router.HandleFunc("/accounts", getAccounts).Methods("GET")
	//Get account by id
	router.HandleFunc("/api/v1/account/{id}", getAccount).Methods("GET")
	//Create new account
	router.HandleFunc("/api/v1/account", createAccount).Methods("POST")
	//Create account balance in account id
	router.HandleFunc("/api/v1/account/{id}/balance", createAccountBalance).Methods("POST")
	//get balance from account id
	router.HandleFunc("/api/v1/account/{id}/balance", getAccountBalance).Methods("GET")
	//Delete account by id
	router.HandleFunc("/api/v1/account/{id}", deleteAccount).Methods("DELETE")

	// start the server
	log.Fatal(http.ListenAndServe(":8080", router))

}
