package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func setupDB() *sql.DB {

	db := connectToDB()

	createTables(db)

	return db
}

func connectToDB() *sql.DB {

	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/todoapp?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}

func createTables(db *sql.DB) {
	{
		query := `
	CREATE TABLE notes (
		id INT AUTO_INCREMENT,
		note TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
);`
		_, err := db.Exec(query)
		if err != nil {
			panic(err.Error())
		}
	}
}