package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
)

func runCpp(input []byte, pathUsersfiles string) (string, error) {
	cppStruct := Cpp{id: uuid.NewString(), userInput: input}
	output, err := getOutput(cppStruct, pathUsersfiles)
	if err != nil {
		return "", err
	}

	return output, nil
}

func runPy(input []byte, pathUsersfiles string) (string, error) {
	pyStruct := Py{id: uuid.NewString(), userInput: input}
	output, err := getOutput(pyStruct, pathUsersfiles)
	if err != nil {
		return "", err
	}

	return output, nil
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
	GenerateFile(pathUserfiles string) (string, error)
	Compile(pathUserfiles string, pathSource string) (string, error)
	Run(pathBinary string) (string, error)
}

type Cpp struct {
	id        string
	userInput []byte
}

func (cpp Cpp) GenerateFile(pathUserfiles string) (string, error) {
	fileName := fmt.Sprintf("%s.cpp", cpp.id)
	path := filepath.Join(pathUserfiles, fileName)
	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %s\n%w", path, err)
	}
	defer f.Close()

	_, err = f.Write(cpp.userInput)
	if err != nil {
		return "", fmt.Errorf("failed to write to file: %s\n%w", path, err)
	}
	f.Sync()

	return path, nil
}

func (cpp Cpp) Compile(pathUserfiles string, pathSource string) (string, error) {
	compiledFileName := fmt.Sprintf("%s.out", cpp.id)
	pathBinary := filepath.Join(pathUserfiles, compiledFileName)

	cmd := exec.Command("g++", pathSource, "-o", pathBinary)
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to compile file: %s\n%w", pathSource, err)
	}

	return pathBinary, nil
}

func (cpp Cpp) Run(pathBinary string) (string, error) {
	out, err := exec.Command(pathBinary).Output()
	if err != nil {
		return "", fmt.Errorf("failed to run binary: %s\n%w", pathBinary, err)
	}

	return string(out), nil
}

type Py struct {
	id        string
	userInput []byte
}

func (py Py) GenerateFile(pathUserfiles string) (string, error) {
	fileName := fmt.Sprintf("%s.py", py.id)
	path := filepath.Join(pathUserfiles, fileName)
	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %s\n%w", path, err)
	}
	defer f.Close()

	_, err = f.Write(py.userInput)
	if err != nil {
		return "", fmt.Errorf("failed to write to file: %s\n%w", path, err)
	}
	f.Sync()

	return path, nil
}

func (py Py) Compile(pathUserfiles string, pathSource string) (string, error) {
	pathBinary := pathSource

	return pathBinary, nil
}

func (py Py) Run(pathBinary string) (string, error) {
	out, err := exec.Command("python3", pathBinary).Output()
	if err != nil {
		return "", fmt.Errorf("failed to run binary: %s\n%w", pathBinary, err)
	}

	return string(out), nil
}

func getOutput(lang Language, pathUserfiles string) (string, error) {
	pathSource, err := lang.GenerateFile(pathUserfiles)
	if err != nil {
		return "", err
	}

	pathBinary, err := lang.Compile(pathUserfiles, pathSource)
	if err != nil {
		return "", err
	}

	out, err := lang.Run(pathBinary)
	if err != nil {
		return "", err
	}

	return out, nil
}
