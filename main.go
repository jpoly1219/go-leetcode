package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jpoly1219/go-leetcode/pkg"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/goleetcode")
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM problems;")
	if err != nil {
		log.Fatal("failed to execute query")
	}
	for res.Next() {
		type problem struct {
			Id          int
			Title       string
			Difficulty  string
			Description string
		}
		var exp problem
		err = res.Scan(&exp.Id, &exp.Title, &exp.Difficulty, &exp.Description)
		if err != nil {
			log.Fatal("failed to scan")
		}
		fmt.Println(exp.Id)
		fmt.Println(exp.Title)
		fmt.Println(exp.Difficulty)
		fmt.Println(exp.Description)
	}

	r := mux.NewRouter()
	r.HandleFunc("/run", pkg.Run)
	r.HandleFunc("/problemsets", pkg.Problemsets)

	log.Fatal(http.ListenAndServe(":8090", r))
}
