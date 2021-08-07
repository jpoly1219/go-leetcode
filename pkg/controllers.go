package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func Run(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint reached")

	userfilesDir := filepath.Join(".", "userfiles/")

	content, _ := ioutil.ReadFile(filepath.Join(userfilesDir, "test.cpp"))
	outCpp, err := runCpp(content, userfilesDir)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(outCpp))

	content, _ = ioutil.ReadFile(filepath.Join(userfilesDir, "test.py"))
	outPy, err := runPy(content, userfilesDir)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(outPy))

}
