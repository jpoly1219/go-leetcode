package main

import (
	"C"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
import "os/exec"

func RunCpp() {
	// receive request with user code inside
	// insert user code into file at a certain place within the template file
	// compile and run
	// read result from result.json
	// save to submissions database. (columns = username, question number, language, code, runtime, result, output)
	// send results and output back as JSON
}

func RunTest(w http.ResponseWriter, r *http.Request) {
	// insert user code
	
	// main_test.go will do the testing of user code and generate an output
	// this will be sent back to the user
	output, err := exec.Command("go", "test").Output()
	if err != nil {
		panic(err.Error())
	}
	w.Write([]byte(output))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run-cpp", RunTest)

	log.Fatal(http.ListenAndServe(":8090", r))
}
