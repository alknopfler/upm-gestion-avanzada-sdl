package main

import (
	"database/sql"
	"log"
	"net/http"
)

func getAccounts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// connect to database
		db := connectDB()
		// close the database connection
		defer closeDB(db)

		// get all accounts from the database

		// write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get all accounts"}`))
	}
	// if the method is not GET, return 405
	w.WriteHeader(http.StatusMethodNotAllowed)
	return

}

func connectDB() *sql.DB {
	// Connect to the database
	conn, err := sql.Open("sqlite3", "/tmp/mybank.db")
	if err != nil {
		log.Println(err)
		return nil
	}
	return conn
}

func closeDB(db *sql.DB) {
	db.Close()
}
