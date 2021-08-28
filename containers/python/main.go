package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RunPy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running python")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run-py", RunPy)

	log.Fatal(http.ListenAndServe(":8090", r))
}
