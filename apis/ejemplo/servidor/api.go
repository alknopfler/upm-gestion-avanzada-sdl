package main

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
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
	}
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
