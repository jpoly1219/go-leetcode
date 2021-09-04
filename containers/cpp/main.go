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

func RunTest(w http.ResponseWriter, r *http.Request) {
	// Insert user code to template.cpp. By inserting code within RunTest, we can control when the test is run.
	// It wouldn't make sense for the test to run when the user code isn't in the file.
	type userCode struct {
		Lang string `json:"lang"`
		Code string `json:"code"`
	}

	var code userCode
	json.NewDecoder(r.Body).Decode(&code)

	templateFile, err := os.OpenFile("template.cpp", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer templateFile.Close()

	codeToInsert := code.Code
	_, err = fmt.Fprintln(templateFile, codeToInsert)
	if err != nil {
		fmt.Println(err)
		templateFile.Close()
		return
	}
	fmt.Println("user code appended successfully")

	// run user code and get any compile or runtime errors using exec.Command().Output()
	cmd := exec.Command("g++", "template.cpp", "-o", "template.out")
	err = cmd.Run()
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

	// return compile or runtime error back to the backend
	if string(out) != "done\n" {
		w.Write(out)
	}

	// read result from result.json
	type resultFile struct {
		Result   string `json:"result"`
		Input    string `json:"input"`
		Expected string `json:"expected"`
		Output   string `json:"output"`
	}
	var result resultFile

	resFile, err := os.ReadFile("result.json")
	// if there are no errors, read the result.json
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(resFile, &result)

	// save to submissions database. (columns = username, question number, language, code, runtime, result, output)
	// send results and output back as JSON
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run-cpp", RunTest)

	log.Fatal(http.ListenAndServe(":8090", r))
}
