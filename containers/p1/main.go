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

// interface and structs/methods definition
type Language interface {
	GenerateFile(templatePath, sourcePath string) error
	CompileAndRun(sourcePath string) (string, error)
}

type Cpp struct {
	Code string
}

func (cpp Cpp) GenerateFile(templatePath, sourcePath string) error {
	lines, err := FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	err = WriteCodeToFile(sourcePath, cpp.Code, lines)
	if err != nil {
		fmt.Println("WriteCodeToFile failed")
		return err
	}
	return nil
}

func (cpp Cpp) CompileAndRun(sourcePath string) (string, error) {
	cmd := exec.Command("g++", "cpp/file.cpp", "-o", "cpp/file.out")
	err := cmd.Run()
	if err != nil {
		fmt.Println("compile failed")
		return "", err
	}

	err = os.Chdir("cpp")
	if err != nil {
		fmt.Println("cd failed")
		return "", err
	}

	out, err := exec.Command("./file.out").Output()
	if err != nil {
		fmt.Println("run failed")
		return "", err
	}
	fmt.Println(string(out))
	return string(out), nil
}

func GetOutput(lang Language, templatePath, sourcePath string) (string, []byte, error) {
	err := lang.GenerateFile(templatePath, sourcePath)
	if err != nil {
		fmt.Println("GenerateFile failed")
		return "", nil, err
	}
	output, err := lang.CompileAndRun(sourcePath)
	if err != nil {
		fmt.Println("CompileAndRun failed")
		return "", nil, err
	}

	if string(output) != "done\n" {
		return output, nil, nil
	}

	result, err := os.ReadFile("result.json")
	if err != nil {
		fmt.Println(err)
		return "", nil, err
	}
	return "", result, nil
}

type Py struct {
	Code string
}

func (py Py) GenerateFile(templatePath, sourcePath string) error {
	lines, err := FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	err = WriteCodeToFile(sourcePath, py.Code, lines)
	if err != nil {
		fmt.Println("WriteCodeToFile failed")
		return err
	}
	return nil
}

func (py Py) CompileAndRun(sourcePath string) (string, error) {
	err := os.Chdir("py")
	if err != nil {
		fmt.Println("cd failed")
		return "", err
	}

	out, err := exec.Command("python3", "file.py").Output()
	if err != nil {
		fmt.Println("run failed")
		return "", err
	}
	fmt.Println(string(out))
	return string(out), nil
}

func RunTest(w http.ResponseWriter, r *http.Request) {
	type userCode struct {
		Pnum int    `json:"pnum"`
		Lang string `json:"lang"`
		Code string `json:"code"`
	}

	var code userCode
	json.NewDecoder(r.Body).Decode(&code)

	var cppCode Cpp
	var pyCode Py
	switch code.Lang {
	case "cpp":
		cppCode.Code = code.Code
		userCodeErr, resultJson, err := GetOutput(cppCode, "cpp/template.cpp", "cpp/file.cpp")
		if err != nil {
			fmt.Println(err)
			return
		}
		if userCodeErr != "" {
			w.Write([]byte(userCodeErr))
			return
		}
		type resultFile struct {
			Result   string `json:"result"`
			Input    string `json:"input"`
			Expected string `json:"expected"`
			Output   string `json:"output"`
		}
		var result resultFile

		json.Unmarshal(resultJson, &result)
		json.NewEncoder(w).Encode(result)
	case "py":
		pyCode.Code = code.Code
		userCodeErr, resultJson, err := GetOutput(cppCode, "py/template.py", "py/file.py")
		if err != nil {
			fmt.Println(err)
			return
		}
		if userCodeErr != "" {
			w.Write([]byte(userCodeErr))
			return
		}
		type resultFile struct {
			Result   string `json:"result"`
			Input    string `json:"input"`
			Expected string `json:"expected"`
			Output   string `json:"output"`
		}
		var result resultFile

		json.Unmarshal(resultJson, &result)
		json.NewEncoder(w).Encode(result)
	}
	/*
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
	*/
	// if there are no errors, read result from result.json

	// save to submissions database. (columns = username, question number, language, code, runtime, result, output)
	// (do this in the backend not the container)
	// send results and output back as JSON
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/run", RunTest)

	log.Fatal(http.ListenAndServe(":8090", r))
}
