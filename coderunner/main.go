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

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jpoly1219/go-leetcode/coderunner/utils"
	_ "github.com/lib/pq"
)

var db *sql.DB

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
	GenerateFile() error
	CompileAndRun() (string, error)
}

type Cpp struct {
	Id       string
	Code     string
	Template string
	RootDir  string
}

func (cpp Cpp) GenerateFile() error {
	// generate template.cpp
	templatePath := filepath.Join(cpp.RootDir, cpp.Id+"-template.cpp")
	templateLines := []byte(cpp.Template)
	err := os.WriteFile(templatePath, templateLines, 0644)
	if err != nil {
		fmt.Println("failed to create template")
	}

	// generate file.cpp
	codeLines, err := utils.FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	sourcePath := filepath.Join(cpp.RootDir, cpp.Id+"-source.cpp")
	err = WriteCodeToFile(sourcePath, cpp.Code, codeLines)
	if err != nil {
		fmt.Println("WriteCodeToFile failed")
		return err
	}
	return nil
}

func (cpp Cpp) CompileAndRun() (string, error) {
	sourcePath := filepath.Join(cpp.RootDir, cpp.Id+"-source.cpp")
	binaryPath := filepath.Join(cpp.RootDir, cpp.Id+"-binary.out")
	cmd := exec.Command("g++", sourcePath, "-o", binaryPath)
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

	runBinaryCommand := "./" + cpp.Id + "-binary.out"
	out, err := exec.Command(runBinaryCommand).Output()
	if err != nil {
		fmt.Println("run failed")
		return "", err
	}

	// fmt.Println(string(out))
	return string(out), nil
}

type Java struct {
	Id       string
	Code     string
	Template string
	RootDir  string
}

func (java Java) GenerateFile() error {
	// generate template.java
	templatePath := filepath.Join(java.RootDir, java.Id+"-template.java")
	templateLines := []byte(java.Template)
	err := os.WriteFile(templatePath, templateLines, 0644)
	if err != nil {
		fmt.Println("failed to create template")
	}

	// generate file.java
	lines, err := utils.FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	sourcePath := filepath.Join(java.RootDir, java.Id+"-source.java")
	err = WriteCodeToFile(sourcePath, java.Code, lines)
	if err != nil {
		fmt.Println("WriteCodeToFile failed")
		return err
	}
	return nil
}

func (java Java) CompileAndRun() (string, error) {
	err := os.Chdir("java")
	if err != nil {
		fmt.Println("cd failed")
		return "", err
	}

	out, err := exec.Command("java", java.Id+"-source.java").Output()
	if err != nil {
		fmt.Println("run failed")
		return "", err
	}

	// fmt.Println(string(out))
	return string(out), nil
}

type Js struct {
	Id       string
	Code     string
	Template string
	RootDir  string
}

func (js Js) GenerateFile() error {
	// generate template.js
	templatePath := filepath.Join(js.RootDir, js.Id+"-template.js")
	templateLines := []byte(js.Template)
	err := os.WriteFile(templatePath, templateLines, 0644)
	if err != nil {
		fmt.Println("failed to create template")
	}

	// generate file.js
	lines, err := utils.FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	sourcePath := filepath.Join(js.RootDir, js.Id+"-source.js")
	err = WriteCodeToFile(sourcePath, js.Code, lines)
	if err != nil {
		fmt.Println("WriteCodeToFile failed")
		return err
	}
	return nil
}

func (js Js) CompileAndRun() (string, error) {
	err := os.Chdir("js")
	if err != nil {
		fmt.Println("cd failed")
		return "", err
	}

	out, err := exec.Command("node", js.Id+"-source.js").Output()
	if err != nil {
		fmt.Println("run failed")
		return "", err
	}

	// fmt.Println(string(out))
	return string(out), nil
}

type Py struct {
	Id       string
	Code     string
	Template string
	RootDir  string
}

func (py Py) GenerateFile() error {
	// generate template.py
	templatePath := filepath.Join(py.RootDir, py.Id+"-template.py")
	templateLines := []byte(py.Template)
	err := os.WriteFile(templatePath, templateLines, 0644)
	if err != nil {
		fmt.Println("failed to create template")
	}

	// generate file.py
	lines, err := utils.FileToLines(templatePath)
	if err != nil {
		fmt.Println("FileToLines failed")
		return err
	}

	sourcePath := filepath.Join(py.RootDir, py.Id+"-source.py")
	err = WriteCodeToFile(sourcePath, py.Code, lines)
	if err != nil {
		fmt.Println("WriteCodeToFile failed")
		return err
	}
	return nil
}

func (py Py) CompileAndRun() (string, error) {
	err := os.Chdir("py")
	if err != nil {
		fmt.Println("cd failed")
		return "", err
	}

	out, err := exec.Command("python3", py.Id+"-source.py").Output()
	if err != nil {
		fmt.Println("run failed")
		return "", err
	}

	// fmt.Println(string(out))
	return string(out), nil
}

func GetOutput(lang Language) (string, []byte, error) {
	err := lang.GenerateFile()
	if err != nil {
		fmt.Println("GenerateFile failed")
		return "", nil, err
	}
	output, err := lang.CompileAndRun()
	if err != nil {
		fmt.Println("CompileAndRun failed")
		return "", nil, err
	}

	if !strings.Contains(output, "test completed\n") {
		return output, nil, nil
	} else {
		resultJson := strings.ReplaceAll(output, "test completed\n", "")
		return "", []byte(resultJson), nil
	}
}

type resultFile struct {
	Username string `json:"username"`
	Slug     string `json:"Slug"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
	Result   string `json:"result"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Output   string `json:"output"`
	// Runtime string `json:"runtime"`
}

func HandleLangs(username, slug, lang, code, template string) (*resultFile, error) {
	var result resultFile
	result.Username = username
	result.Slug = slug
	result.Lang = lang
	result.Code = code

	var userCode Language
	// fmt.Println(lang)
	switch lang {
	case "cpp":
		userCode = Cpp{Id: uuid.NewString(), Code: code, Template: template, RootDir: "cpp/"}
	case "java":
		userCode = Java{Id: uuid.NewString(), Code: code, Template: template, RootDir: "java/"}
	case "js":
		userCode = Js{Id: uuid.NewString(), Code: code, Template: template, RootDir: "js/"}
	case "py":
		userCode = Py{Id: uuid.NewString(), Code: code, Template: template, RootDir: "py/"}
	}

	// fmt.Println("now running GetOutput: userCode: ", userCode)
	userCodeErr, resultJson, err := GetOutput(userCode)
	if err != nil {
		fmt.Println(err)
		result.Result = "wrong"
		result.Input = ""
		result.Expected = ""
		result.Output = fmt.Sprintf("%s", err)
		return &result, nil
	}
	if userCodeErr != "" {
		result.Result = "wrong"
		result.Input = ""
		result.Expected = ""
		result.Output = userCodeErr
		return &result, nil
	}

	json.Unmarshal(resultJson, &result)
	fmt.Println(result)
	return &result, nil
}

func RunTest(w http.ResponseWriter, r *http.Request) {
	// may be unnecessary when deploying the app using docker compose
	utils.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	type userCode struct {
		Username string `json:"username"`
		Slug     string `json:"slug"`
		Lang     string `json:"lang"`
		Code     string `json:"code"`
	}

	var code userCode
	json.NewDecoder(r.Body).Decode(&code)
	// fmt.Println("RunTest() reached: ", code.Username, code.Slug, code.Lang, code.Code)

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	if filepath.Base(cwd) != "coderunner" {
		err = os.Chdir("..")
		if err != nil {
			fmt.Println("cd failed")
			fmt.Println(err)
			return
		}
	}

	var testcase string
	err = db.QueryRow(
		"SELECT testcase FROM testcases WHERE slug = $1;",
		code.Slug,
	).Scan(&testcase)
	if err != nil {
		fmt.Println("failed to query and scan to testcase: ", err)
		return
	}

	testcasePath := "./testcase-" + code.Slug + ".json"
	err = os.WriteFile(testcasePath, []byte(testcase), 0644)
	if err != nil {
		fmt.Println("failed to create testcase.json")
	}

	var template string
	err = db.QueryRow(
		"SELECT template FROM templates WHERE slug = $1 AND lang = $2;",
		code.Slug, code.Lang,
	).Scan(&template)
	if err != nil {
		fmt.Println("failed to query and scan to template: ", err)
		return
	}

	result, err := HandleLangs(code.Username, code.Slug, code.Lang, code.Code, template)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*result)

	// save attempt to database
	_, err = db.Exec(
		"INSERT INTO attempts (username, slug, lang, code, result, output) VALUES ($1, $2, $3, $4, $5, $6);",
		result.Username, result.Slug, result.Lang, result.Code, result.Result, result.Output,
	)
	if err != nil {
		fmt.Println("failed to insert attempt: ", err)
		return
	}
	json.NewEncoder(w).Encode(*result)
}

func main() {
	// establish database connection
	var err error
	const (
		host     = "jpoly1219devbox.xyz"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "goleetcode"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
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
