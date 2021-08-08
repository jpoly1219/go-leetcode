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

	pathUserfiles := filepath.Join(".", "userfiles", "testuserfiles")

	content, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test.cpp"))
	outCpp, err := runCpp(content, pathUserfiles)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(outCpp))

	content, _ = ioutil.ReadFile(filepath.Join(pathUserfiles, "test.java"))
	outJava, err := runJava(content, pathUserfiles)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(outJava))

	content, _ = ioutil.ReadFile(filepath.Join(pathUserfiles, "test.js"))
	outJs, err := runJs(content, pathUserfiles)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(outJs))

	content, _ = ioutil.ReadFile(filepath.Join(pathUserfiles, "test.py"))
	outPy, err := runPy(content, pathUserfiles)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(outPy))
}
