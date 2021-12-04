package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

// outdated and for testing purposes only!
func fileGen(language string, userInput string, pathUserfiles string) string {
	var out string
	var err error
	switch language {
	case "C++":
		out, err = RunCpp([]byte(userInput), pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "Java":
		out, err = RunJava([]byte(userInput), pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "Javascript":
		out, err = RunJs([]byte(userInput), pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	case "Python":
		out, err = RunPy([]byte(userInput), pathUserfiles)
		if err != nil {
			log.Fatal(err)
		}
		// w.Write
		fmt.Println(out)
	}
	return out
}

// outdated and for testing purposes only!
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
}

func Problemsets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reached Problemsets")
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
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

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
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var input userCode
	json.NewDecoder(r.Body).Decode(&input)

	fmt.Println("CheckProblem() reached: ", input.Username, input.Slug, input.Lang, input.Code)

	// use Docker SDK to run a container to run user code safely inside a sandbox
	// then send a POST request which contains several fields to the container
	postBody, _ := json.Marshal(map[string]string{
		"username": input.Username,
		"slug":     input.Slug,
		"lang":     input.Lang,
		"code":     input.Code,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://jpoly1219devbox.xyz:8091/run", "application/json", responseBody)
	if err != nil {
		fmt.Println("POST request failed: ", err)
		return
	}
	defer resp.Body.Close()

	var resFromContainer result
	json.NewDecoder(resp.Body).Decode(&resFromContainer)
	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(resFromContainer)
}

func Submissions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reached submissions")
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var userSubmission submission
	json.NewDecoder(r.Body).Decode(&userSubmission)
	fmt.Println(userSubmission.Username, userSubmission.Slug)

	var prevSubmissions = make([]result, 0)
	results, err := Db.Query(
		"SELECT username, slug, lang, code, result, output FROM attempts WHERE username = $1 AND slug = $2;",
		userSubmission.Username, userSubmission.Slug,
	)
	if err != nil {
		log.Fatal("failed to query attempts")
	}
	for results.Next() {
		var prevSubmission result
		err = results.Scan(
			&prevSubmission.Username, &prevSubmission.Slug, &prevSubmission.Lang,
			&prevSubmission.Code, &prevSubmission.Result, &prevSubmission.Output,
		)
		if err != nil {
			log.Fatal("failed to scan")
		}
		prevSubmissions = append(prevSubmissions, prevSubmission)
	}

	fmt.Println("hi", prevSubmissions)
	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(prevSubmissions)
}

func Solutions(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["slug"]

	var s solution
	err := Db.QueryRow("SELECT solution FROM solutions WHERE slug = $1;", keys).Scan(&s.Solution)
	if err != nil {
		log.Fatal("failed to execute query", err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(s)
}

func Discussions(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["slug"]

	var discussions = make([]discussion, 0)
	results, err := Db.Query("SELECT * FROM discussions WHERE slug = $1;", keys)
	if err != nil {
		log.Fatal("failed to execute query", err)
		return
	}
	for results.Next() {
		var d discussion
		err = results.Scan(&d.Id, &d.Author, &d.Slug, &d.Title, &d.Description, &d.Created)
		if err != nil {
			log.Fatal("failed to scan", err)
		}
		discussions = append(discussions, d)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(discussions)
}

func Comments(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["discussionId"]

	var comments = make([]comment, 0)
	results, err := Db.Query("SELECT * FROM comments WHERE discussion_id = $1;", keys)
	if err != nil {
		log.Fatal("failed to execute query, err")
	}
	for results.Next() {
		var c comment
		err = results.Scan(&c.Id, &c.Author, &c.DiscussionId, &c.Description, &c.Created)
		if err != nil {
			log.Fatal("failed to scan", err)
		}
		comments = append(comments, c)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	json.NewEncoder(w).Encode(comments)
}

func NewComment(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["discussionId"]

	var newComment comment
	json.NewDecoder(r.Body).Decode(&newComment)
	fmt.Println("new comment: ", newComment)

	err := Db.QueryRow(
		"INSERT INTO comments (author, discussion_id, description) VALUES ($1, $2, $3) RETURNING *;",
		&newComment.Author, keys, &newComment.Description,
	).Scan(&newComment.Id, &newComment.Author, &newComment.DiscussionId, &newComment.Description, &newComment.Created)
	if err != nil {
		fmt.Println("failed to insert comment: ", err)
		return
	}

	json.NewEncoder(w).Encode(&newComment)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["username"]

	var u user
	err := Db.QueryRow(
		"SELECT username, fullname, email, profile_pic FROM users WHERE username = $1;",
		keys,
	).Scan(&u.Username, &u.Fullname, &u.Email, &u.ProfilePic)
	if err != nil {
		fmt.Println("failed to query user: ", err)
		return
	}

	fmt.Println(u)

	json.NewEncoder(w).Encode(&u)
}

func NewDiscussion(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["slug"]

	var newDiscussion discussion
	json.NewDecoder(r.Body).Decode(&newDiscussion)
	fmt.Println("new discussion: ", newDiscussion)

	err := Db.QueryRow(
		"INSERT INTO comments (author, slug, title, description) VALUES ($1, $2, $3, $4) RETURNING *;",
		&newDiscussion.Author, keys, &newDiscussion.Title, &newDiscussion.Description,
	).Scan(&newDiscussion.Id, &newDiscussion.Author, &newDiscussion.Slug, &newDiscussion.Title, &newDiscussion.Description, &newDiscussion.Created)
	if err != nil {
		fmt.Println("failed to insert comment: ", err)
		return
	}

	json.NewEncoder(w).Encode(&newDiscussion)
}
