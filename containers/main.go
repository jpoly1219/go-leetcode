package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

var db *sql.DB

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

	codeLines := strings.Split(code, "\n")

	writer := bufio.NewWriter(f)
	for _, line := range lines {
		if strings.Contains(line, "insert Solution class here") {
			for _, codeLine := range codeLines {
				_, _ = writer.WriteString(codeLine + "\n")
			}
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
	Code     string
	Template string
}

func (cpp Cpp) GenerateFile(templatePath, sourcePath string) error {
	// generate template.cpp
	templateLines := []byte(cpp.Template)
	err := os.WriteFile(templatePath, templateLines, 0644)
	if err != nil {
		fmt.Println("failed to create template")
	}

	// generate file.cpp
	codeLines, err := FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	err = WriteCodeToFile(sourcePath, cpp.Code, codeLines)
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

	err = os.Chdir("..")
	if err != nil {
		fmt.Println("cd failed")
	}

	fmt.Println(string(out))
	return string(out), nil
}

type Java struct {
	Code     string
	Template string
}

func (java Java) GenerateFile(templatePath, sourcePath string) error {
	// generate template.java
	templateLines := []byte(java.Template)
	err := os.WriteFile(templatePath, templateLines, 0644)
	if err != nil {
		fmt.Println("failed to create template")
	}

	// generate file.java
	lines, err := FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	err = WriteCodeToFile(sourcePath, java.Code, lines)
	if err != nil {
		fmt.Println("WriteCodeToFile failed")
		return err
	}
	return nil
}

func (java Java) CompileAndRun(sourcePath string) (string, error) {
	err := os.Chdir("js")
	if err != nil {
		fmt.Println("cd failed")
		return "", err
	}

	out, err := exec.Command("java", "file.java").Output()
	if err != nil {
		fmt.Println("run failed")
		return "", err
	}
	fmt.Println(string(out))
	return string(out), nil
}

type Js struct {
	Code     string
	Template string
}

func (js Js) GenerateFile(templatePath, sourcePath string) error {
	// generate template.js
	templateLines := []byte(js.Template)
	err := os.WriteFile(templatePath, templateLines, 0644)
	if err != nil {
		fmt.Println("failed to create template")
	}

	// generate file.js
	lines, err := FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	err = WriteCodeToFile(sourcePath, js.Code, lines)
	if err != nil {
		fmt.Println("WriteCodeToFile failed")
		return err
	}
	return nil
}

func (js Js) CompileAndRun(sourcePath string) (string, error) {
	err := os.Chdir("js")
	if err != nil {
		fmt.Println("cd failed")
		return "", err
	}

	out, err := exec.Command("node", "file.js").Output()
	if err != nil {
		fmt.Println("run failed")
		return "", err
	}
	fmt.Println(string(out))
	return string(out), nil
}

type Py struct {
	Code     string
	Template string
}

func (py Py) GenerateFile(templatePath, sourcePath string) error {
	// generate template.py
	templateLines := []byte(py.Template)
	err := os.WriteFile(templatePath, templateLines, 0644)
	if err != nil {
		fmt.Println("failed to create template")
	}

	// generate file.py
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

type resultFile struct {
	Username string `json:"username"`
	Pnum     int    `json:"pnum"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
	Result   string `json:"result"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Output   string `json:"output"`
	// Runtime string `json:"runtime"`
}

func HandleLangs(pnum int, username, code, lang, template string) (*resultFile, error) {
	var result resultFile
	result.Username = username
	result.Pnum = pnum
	result.Lang = lang
	result.Code = code

	switch lang {
	case "C++":
		/*
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if filepath.Base(cwd) == "cpp" {
				err := os.Chdir("..")
				if err != nil {
					fmt.Println("cd failed")
				}
			}
		*/
		cppCode := Cpp{Code: code, Template: template}
		userCodeErr, resultJson, err := GetOutput(cppCode, "cpp/template.cpp", "cpp/file.cpp")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if userCodeErr != "" {
			result.Result = "wrong"
			result.Input = ""
			result.Expected = ""
			result.Output = userCodeErr
			return &result, nil
		}

		json.Unmarshal(resultJson, &result)
	case "Java":
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if filepath.Base(cwd) == "java" {
			err := os.Chdir("..")
			if err != nil {
				fmt.Println("cd failed")
			}
		}

		javaCode := Java{Code: code, Template: template}
		userCodeErr, resultJson, err := GetOutput(javaCode, "java/template.java", "java/file.java")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if userCodeErr != "" {
			result.Result = "wrong"
			result.Input = ""
			result.Expected = ""
			result.Output = userCodeErr
			return &result, nil
		}

		json.Unmarshal(resultJson, &result)
	case "Javascript":
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if filepath.Base(cwd) == "js" {
			err := os.Chdir("..")
			if err != nil {
				fmt.Println("cd failed")
			}
		}

		jsCode := Js{Code: code, Template: template}
		userCodeErr, resultJson, err := GetOutput(jsCode, "js/template.js", "js/file.js")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if userCodeErr != "" {
			result.Result = "wrong"
			result.Input = ""
			result.Expected = ""
			result.Output = userCodeErr
			return &result, nil
		}

		json.Unmarshal(resultJson, &result)
	case "Python":
		fmt.Println("python detected")

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if filepath.Base(cwd) == "py" {
			err := os.Chdir("..")
			if err != nil {
				fmt.Println("cd failed")
			}
		}

		pyCode := Py{Code: code, Template: template}
		userCodeErr, resultJson, err := GetOutput(pyCode, "py/template.py", "py/file.py")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if userCodeErr != "" {
			result.Result = "wrong"
			result.Input = ""
			result.Expected = ""
			result.Output = userCodeErr
			return &result, nil
		}

		json.Unmarshal(resultJson, &result)
	}
	fmt.Println(result)
	return &result, nil
}

func RunTest(w http.ResponseWriter, r *http.Request) {
	type userCode struct {
		Username string `json:"username"`
		Pnum     int    `json:"pnum"`
		Lang     string `json:"lang"`
		Code     string `json:"code"`
	}

	var code userCode
	json.NewDecoder(r.Body).Decode(&code)
	fmt.Println("RunTest() reached: ", code.Username, code.Pnum, code.Lang, code.Code)

	queryResult, err := db.Query(
		"SELECT template, testcase FROM tests WHERE lang = ? AND problem_id = ?",
		code.Lang, code.Pnum,
	)
	if err != nil {
		fmt.Println("failed to execute query: ", err)
		return
	}
	var template, testcase string
	for queryResult.Next() {
		err = queryResult.Scan(&template, &testcase)
		if err != nil {
			log.Fatal("failed to scan")
		}
	}

	result, err := HandleLangs(code.Pnum, code.Username, code.Code, code.Lang, template)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*result)
	json.NewEncoder(w).Encode(*result)

	// save to submissions database. (columns = username, question number, language, code, runtime, result, output)
}

func main() {
	// establish database connection
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/goleetcode")
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/run", RunTest)

	// each container will have the exact same domain, so how will the backend distinguish between the containers?
	// plus, if all containers use port 8091, it would be even more confusing...
	// may need a container orchestration tool
	log.Fatal(http.ListenAndServe(":8091", r))
}
