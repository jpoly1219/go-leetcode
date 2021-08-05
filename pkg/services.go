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
