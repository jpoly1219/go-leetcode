package pkg

import (
	"fmt"
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

	cmd := exec.Command("touch", pathToFile)
	fmt.Println("creating file")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmdInsertInput := fmt.Sprintf("%s > %s.cpp", string(input), fileId)
	cmd = exec.Command(cmdInsertInput)
	fmt.Println("inserting input")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	cmdCompileAndRun := fmt.Sprintf("g++ %s -o %s; ./%s", pathToFile, fileId, fileId)
	fmt.Println("compiling and running")
	out, err := exec.Command(cmdCompileAndRun).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("output:\n", string(out))
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
