package main

import (
	"C"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunCpp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running cpp")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run-cpp", RunCpp)

	log.Fatal(http.ListenAndServe(":8090", r))
}
