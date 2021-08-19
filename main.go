package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jpoly1219/go-leetcode/pkg"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	pkg.Db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/goleetcode")
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	defer pkg.Db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/run", pkg.Run)
	r.HandleFunc("/problemsets", pkg.Problemsets)

	log.Fatal(http.ListenAndServe(":8090", r))
}
