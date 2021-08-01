package main

import (
	"log"
	"net/http"

	"github.com/jpoly1219/go-leetcode/pkg"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/run", pkg.SayHello)

	log.Fatal(http.ListenAndServe(":8090", r))
}
