package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

func FileToLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	return LinesFromFile(f)
}

func LinesFromFile(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func WriteCodeToFile(filePath, code string, lines []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, line := range lines {
		if line == "// insert Solution class here" {
			_, _ = writer.WriteString(code + "\n")
		}
		_, _ = writer.WriteString(line + "\n")
	}
	writer.Flush()
	return nil
}

func RunTest(w http.ResponseWriter, r *http.Request) {
	type userCode struct {
		Lang string `json:"lang"`
		Code string `json:"code"`
	}

	var code userCode
	json.NewDecoder(r.Body).Decode(&code)

	// read template.cpp, save each line as slice, insert user code into it, then write the slice
	lines, err := FileToLines("template.cpp")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = WriteCodeToFile("file.cpp", code.Code, lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user code appended successfully")

	// run user code and get any compile or runtime errors using exec.Command().Output()
	cmd := exec.Command("g++", "file.cpp", "-o", "file.out")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := exec.Command("./file.out").Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

	// return compile or runtime error back to the backend
	if string(out) != "done\n" {
		w.Write(out)
	}

	// if there are no errors, read result from result.json
	type resultFile struct {
		Result   string `json:"result"`
		Input    string `json:"input"`
		Expected string `json:"expected"`
		Output   string `json:"output"`
	}
	var result resultFile

	resFile, err := os.ReadFile("result.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(resFile, &result)

	// save to submissions database. (columns = username, question number, language, code, runtime, result, output)
	// (do this in the backend not the container)
	// send results and output back as JSON
	json.NewEncoder(w).Encode(result)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run-cpp", RunTest)

	log.Fatal(http.ListenAndServe(":8090", r))
}
