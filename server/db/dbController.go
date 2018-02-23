package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect() {
	db, err := sql.Open("postgres", "postgres://nbkmphdo:PCTacSE0ocbRGxQloL4yq9g8-XRBqxlN@baasu.db.elephantsql.com:5432/nbkmphdo")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database successfully connected!")
}
