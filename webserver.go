package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db := setupDB()

	r := mux.NewRouter()

	noteHandler := func(w http.ResponseWriter, r *http.Request) {
		createNoteHandler(w, r, db)
	}

	r.HandleFunc("/notes", noteHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Fprintf(w, "You've requested to create a note\n")
	note := r.FormValue("note")
	created_at := time.Now()

	result, err := db.Exec(`INSERT INTO notes (note, created_at) VALUES (?, ?)`, note, created_at)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
}