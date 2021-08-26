package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func fileGen(language string, userInput string, pathUserfiles string) {
	switch language {
	case "C++":
		out, err := RunCpp([]byte(userInput), pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "Java":
		out, err := RunJava([]byte(userInput), pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "Javascript":
		out, err := RunJs([]byte(userInput), pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "Python":
		out, err := RunPy([]byte(userInput), pathUserfiles)
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
		go fileGen("test.cpp", "hi", pathUserfiles)
		go fileGen("test.java", "hi", pathUserfiles)
		go fileGen("test.js", "hi", pathUserfiles)
		go fileGen("test.py", "hi", pathUserfiles)
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

	results, err := Db.Query("SELECT * FROM problems;")
	if err != nil {
		log.Fatal("failed to execute query", err)
	}
	for results.Next() {
		var p problem
		err = results.Scan(&p.Id, &p.Title, &p.Slug, &p.Difficulty, &p.Description, &p.Created)
		if err != nil {
			log.Fatal("failed to scan", err)
		}
		problems = append(problems, p)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(problems)
}

func ReturnProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["slug"]

	var p problem
	results, err := Db.Query("SELECT * FROM problems WHERE slug = $1;", keys)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}
	for results.Next() {
		err = results.Scan(&p.Id, &p.Title, &p.Slug, &p.Difficulty, &p.Description, &p.Created)
		if err != nil {
			log.Fatal("failed to scan")
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(p)
}

func CheckProblem(w http.ResponseWriter, r *http.Request) {
	var input userCode
	json.NewDecoder(r.Body).Decode(&input)

	fmt.Println(input.Lang, input.Code)

	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	w.Write([]byte("Code accepted!"))
}
