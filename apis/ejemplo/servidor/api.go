package main

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func getAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// connect to database
		db := connectDB()
		// close the database connection
		defer closeDB(db)

		// get all accounts from the database
		rows, err := db.Query("SELECT * FROM account")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
			return
		}
		defer rows.Close()
		// Parse row into Activity struct
		accounts := ListAccount{}
		for rows.Next() {
			i := Account{}
			err = rows.Scan(&i.ID, &i.Name, &i.Balance, &i.Currency, &i.Owner, &i.CreatedAt, &i.UpdatedAt, &i.DeletedAt)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			accounts.List = append(accounts.List, i)
		}
		// write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		bodyResponse, _ := json.Marshal(accounts)
		w.Write(bodyResponse)
	}
	// if the method is not GET, return 405
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && mux.Vars(r)["id"] != "" {
		db := connectDB()
		// close the database connection
		defer closeDB(db)

		// get all accounts from the database
		row := db.QueryRow("SELECT * FROM account WHERE id=?", mux.Vars(r)["id"])
		result := Account{}
		var err error
		if err = row.Scan(&result.ID,
			&result.Name,
			&result.Balance,
			&result.Currency,
			&result.Owner,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt); err == sql.ErrNoRows {

			log.Printf("Id %v not found", mux.Vars(r)["id"])
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		bodyResponse, _ := json.Marshal(result)
		w.Write(bodyResponse)
	}
	// if the method is not GET, return 405
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		db := connectDB()
		// close the database connection
		defer closeDB(db)

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		inputAccount := InputAccount{}
		err = json.Unmarshal(b, &inputAccount)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = db.Exec("INSERT INTO account VALUES(?,?,?,?,?,?,?,?);",
			newUUID(),
			inputAccount.Name,
			inputAccount.Balance,
			inputAccount.Currency,
			inputAccount.Owner,
			time.Now().Format("2006-01-02 15:04:05"),
			"",
			"")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func createAccountBalance(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		db := connectDB()
		// close the database connection
		defer closeDB(db)

		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		inputBalance := Balance{}
		err = json.Unmarshal(b, &inputBalance)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = db.Exec("UPDATE account SET balance = ? WHERE id = ?", inputBalance.Balance, mux.Vars(r)["id"])
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func getAccountBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && mux.Vars(r)["id"] != "" {
		db := connectDB()
		// close the database connection
		defer closeDB(db)

		// get all accounts from the database
		row := db.QueryRow("SELECT balance FROM account WHERE id=?", mux.Vars(r)["id"])
		result := Balance{}
		var err error
		if err = row.Scan(&result.Balance); err == sql.ErrNoRows {

			log.Printf("Id %v not found", mux.Vars(r)["id"])
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		bodyResponse, _ := json.Marshal(result)
		w.Write(bodyResponse)
		return
	}
	// if the method is not GET, return 405
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" && mux.Vars(r)["id"] != "" {
		db := connectDB()
		// close the database connection
		defer closeDB(db)

		// get all accounts from the database
		_, err := db.Exec("DELETE FROM account WHERE id=?", mux.Vars(r)["id"])
		if err != nil {
			log.Printf("Id %v not found", mux.Vars(r)["id"])
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
	// if the method is not DELETE, return 405
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
	if _, err := conn.Exec(schemaDB); err != nil {
		return nil
	}
	return conn
}

func closeDB(db *sql.DB) {
	db.Close()
}

func newUUID() string {
	u, _ := uuid.NewUUID()
	return u.String()
}

func destroy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	go func() {
		time.Sleep(1 * time.Minute)
		os.Exit(1)
	}()
}
