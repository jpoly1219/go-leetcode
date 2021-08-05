package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func LoadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

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

// interpreted languages don't need this because it basically runs without needing additional steps
type CompiledLang interface {
	Compile(fileId, path string) string // returns output or error
}

// all languages will need to be run at some point
type Language interface {
	Run() string // returns output or error
}

// all languages should be able to generate a file with their respective user inputs and extensions
type FileGenerator interface {
	Create() error    // touches file and inserts user input
	FindPath() string // returns a file path string
}

type Cpp struct {
	id        string
	userInput []byte
}
type Py struct {
	id        string
	userInput []byte
}

func (cpp Cpp) Compile(fileId, path string) error {
	cmd := exec.Command("g++", path, "-o", fileId)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
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
