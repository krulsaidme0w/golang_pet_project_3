package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"krulsaidme0w/library/pkg/scripts"
)

const (
	count = 1000000
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "user", "password", "library"))
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1000)
	defer db.Close()

	err = scripts.FillLibraryDB(db, 1000000)
	if err != nil {
		log.Fatal(err)
	}
}
