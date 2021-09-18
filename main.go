package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/jpoly1219/go-leetcode/pkg"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file.")
	}

	pkg.Db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/goleetcode")
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	defer pkg.Db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/run", pkg.Run)
	r.HandleFunc("/problemsets", pkg.Problemsets)

	solveR := r.PathPrefix("/solve").Subrouter()
	solveR.HandleFunc("/{slug}", pkg.ReturnProblem)

	checkR := r.PathPrefix("/check").Subrouter()
	checkR.HandleFunc("/{slug}", pkg.CheckProblem)

	log.Fatal(http.ListenAndServe(":8090", r))
}
