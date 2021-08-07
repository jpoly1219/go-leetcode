package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
)

func runCpp(input []byte, userfileDir string) {
	fileId := uuid.New()
	fileName := fmt.Sprintf("%s.cpp", fileId)
	pathToFile := filepath.Join(userfileDir, fileName)
	fmt.Println(pathToFile)

	cmd := exec.Command("touch", pathToFile)
	fmt.Println("creating file")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(pathToFile, input, 0644)
	if err != nil {
		log.Fatal(err)
	}

	compiledFileName := fmt.Sprintf("%s.out", fileId)
	pathToCompiledFile := filepath.Join(userfileDir, compiledFileName)
	cmd = exec.Command("g++", pathToFile, "-o", pathToCompiledFile)
	fmt.Println("compiling and running")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command(pathToCompiledFile).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("output:\n", string(out))
}

func runPy(input []byte, userfileDir string) {

}

/*
implementing interfaces
the following steps apply to all languages:
accept user input > create file with the user input > compile > run > return output as string or error
steps that are same across all languages: 1, 5
steps that are different across languages: 2, 3, 4
- 2: file extensions are different
- 3: some languages need to be compiled, but some don't. For interpreted languages, compile() may as well be a pass statement
- 4: different languages require different tools to run their files, so the command will differ across languages
*/
type Language interface {
	GenerateFile() error
	Compile() error
	Run() (string, error)
}

type Cpp struct {
	id        string
	userInput []byte
}

func (cpp Cpp) GenerateFile() error {
	return nil
}

func (cpp Cpp) Compile() error {
	return nil
}

func (cpp Cpp) Run() error {
	return nil
}

type Py struct {
	id        string
	userInput []byte
}

func (py Py) GenerateFile() error {
	return nil
}

func (py Py) Compile() error {
	return nil
}

func (py Py) Run() error {
	return nil
}

func parseOutput(lang Language) (string, error) {
	err := lang.Compile()
	if err != nil {
		return "error", err
	}
	output, err := lang.Run()
	if err != nil {
		return "error", err
	}
	return output, nil
}

/*
func PrintHello(num int, out chan string) {
	defer close(out)
	for i := 0; i < num; i++ {
		out <- "hello world!\n"
		time.Sleep(time.Second)
	}
}

func PrintBye(num int, out chan string) {
	defer close(out)
	for i := 0; i < num; i++ {
		out <- "goodbye!\n"
		time.Sleep(time.Second)
	}
}
*/
