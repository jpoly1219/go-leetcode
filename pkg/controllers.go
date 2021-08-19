package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func fileGen(testfile string, pathUserfiles string) {
	content, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, testfile))
	switch testfile {
	case "test.cpp":
		out, err := RunCpp(content, pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "test.java":
		out, err := RunJava(content, pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "test.js":
		out, err := RunJs(content, pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "test.py":
		out, err := RunPy(content, pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	}
}

func Run(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint reached")

	pathUserfiles := filepath.Join(".", "userfiles", "testuserfiles")

	// reading from test.* files is only temporary; later on these will be actual user inputs sent over from the frontend

	for i := 0; i < 10; i++ {
		go fileGen("test.cpp", pathUserfiles)
		go fileGen("test.java", pathUserfiles)
		go fileGen("test.js", pathUserfiles)
		go fileGen("test.py", pathUserfiles)
	}
	w.Write([]byte("done\n"))
	/*
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
	*/
}

func Problemsets(w http.ResponseWriter, r *http.Request) {
	var problems = make([]problem, 0)

	result, err := Db.Query("SELECT * FROM problems;")
	if err != nil {
		log.Fatal("failed to execute query")
	}
	for result.Next() {
		var p problem
		err = result.Scan(&p.Id, &p.Title, &p.Difficulty, &p.Description)
		if err != nil {
			log.Fatal("failed to scan")
		}
		problems = append(problems, p)
		fmt.Println(p.Id)
		fmt.Println(p.Title)
		fmt.Println(p.Difficulty)
		fmt.Println(p.Description)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(problems)
}
