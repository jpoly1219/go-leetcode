package lang

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/jpoly1219/go-leetcode/coderunner/models"
	"github.com/jpoly1219/go-leetcode/coderunner/utils"
)

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
	err = utils.WriteCodeToFile(sourcePath, cpp.Code, codeLines)
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

	cwd, _ := os.Getwd()
	fmt.Printf("cwd from car: %s", cwd)
	if filepath.Base(cwd) != "coderunner" {
		err = os.Chdir("..")
		if err != nil {
			fmt.Println("cd failed")
			fmt.Println(err)
		}
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
	err = utils.WriteCodeToFile(sourcePath, java.Code, lines)
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
	err = utils.WriteCodeToFile(sourcePath, js.Code, lines)
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
	err = utils.WriteCodeToFile(sourcePath, py.Code, lines)
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
	fmt.Printf("\noutput: %s\n", output)

	if !strings.Contains(output, "test completed\n") {
		return output, nil, nil
	} else {
		resultJson := strings.ReplaceAll(output, "test completed\n", "")
		return "", []byte(resultJson), nil
	}
}

func HandleLangs(username, slug, lang, code, template string) (*models.ResultFile, error) {
	var result models.ResultFile
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
	_, resultJson, err := GetOutput(userCode)
	type CodeResult struct {
		Result   string
		Input    string
		Expected string
		Output   string
	}
	var cr CodeResult
	json.Unmarshal(resultJson, &cr)

	if err != nil {
		// "", "", err
		fmt.Println(err)
		result.Result = "wrong"
		result.Input = ""
		result.Expected = ""
		result.Output = fmt.Sprintf("%s", err)
		return &result, nil
	}
	// if userCodeErr != "" {
	// 	//
	// 	result.Result = "wrong"
	// 	result.Input = ""
	// 	result.Expected = ""
	// 	result.Output = userCodeErr
	// 	return &result, nil
	// }

	json.Unmarshal(resultJson, &result)
	fmt.Println(result)
	return &result, nil
}
