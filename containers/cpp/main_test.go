package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestRunCpp(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/goleetcode")
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	defer db.Close()

}
