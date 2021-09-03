package main

import (
	"log"
	"net/http"
	"os/exec"

	"encoding/json"
	"fmt"
	"os"

	"github.com/gorilla/mux"
)

func RunCpp() {
	// receive request with user code inside (RunTest)
	// insert user code into file at a certain place within the template file (RunTest)
	// compile and run
	cmd := exec.Command("g++", "template.cpp", "-o", "template.out")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := exec.Command("template.out").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

	// read result from result.json
	type resultFile struct {
		Result   string `json:"result"`
		Input    string `json:"input"`
		Expected string `json:"expected"`
		Output   string `json:"output"`
	}
	var result resultFile

	f, err := os.ReadFile("result.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(f, &result)
	// save to submissions database. (columns = username, question number, language, code, runtime, result, output)
	// send results and output back as JSON
}

func RunTest(w http.ResponseWriter, r *http.Request) {
	// Insert user code to template.cpp. By inserting code within RunTest, we can control when the test is run.
	// It wouldn't make sense for the test to run when the user code isn't in the file.
	type userCode struct {
		Lang string `json:"lang"`
		Code string `json:"code"`
	}

	var code userCode
	json.NewDecoder(r.Body).Decode(&code)

	f, err := os.OpenFile("template.cpp", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	codeToInsert := code.Code
	_, err = fmt.Fprintln(f, codeToInsert)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println("user code appended successfully")
	/*
		// main_test.go will do the testing of user code and generate an output
		// this will be sent back to the user
		output, err := exec.Command("go", "test").Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Write([]byte(output))
	*/
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run-cpp", RunTest)

	log.Fatal(http.ListenAndServe(":8090", r))
}
