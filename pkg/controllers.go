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
	type problem struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Difficulty  string `json:"difficulty"`
		Description string `json:"description"`
	}

	problem1 := problem{
		1, "Problem 1", "easy", "This is problem 1.",
	}
	problem2 := problem{
		1, "Problem 2", "medium", "This is problem 2.",
	}
	problem3 := problem{
		1, "Problem 3", "easy", "This is problem 3.",
	}
	problem4 := problem{
		1, "Problem 4", "hard", "This is problem 4.",
	}
	problem5 := problem{
		1, "Problem 5", "medium", "This is problem 5.",
	}

	problems := []problem{
		problem1, problem2, problem3, problem4, problem5,
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(problems)
}
