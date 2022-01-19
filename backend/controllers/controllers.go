package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jpoly1219/go-leetcode/backend/auth"
	"github.com/jpoly1219/go-leetcode/backend/models"

	"github.com/gorilla/mux"
)

func Problemsets(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	type username struct {
		Username string `json:"username"`
	}

	var uname username
	json.NewDecoder(r.Body).Decode(&uname)

	fmt.Println("reached Problemsets")
	var problems = make([]models.ProblemAndResult, 0)
	var result sql.NullString

	results, err := models.Db.Query(
		"SELECT DISTINCT problems.id, title, problems.slug, difficulty, result FROM problems LEFT JOIN attempts ON problems.slug = attempts.slug AND username = $1 AND result = 'OK' ORDER BY title;",
		uname.Username,
	)
	if err != nil {
		log.Println("failed to execute query", err)
	}
	for results.Next() {
		var p models.ProblemAndResult
		err = results.Scan(&p.Id, &p.Title, &p.Slug, &p.Difficulty, &result)
		if err != nil {
			log.Println("failed to scan", err)
		}

		if result.Valid {
			p.Result = result.String
		} else {
			p.Result = "-"
		}
		problems = append(problems, p)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(problems)
}

func FilterProblemsets(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var f models.Filter
	json.NewDecoder(r.Body).Decode(&f)

	if f.Difficulty == "all" {
		Problemsets(w, r)
	} else {
		var problems = make([]models.ProblemAndResult, 0)
		var result sql.NullString

		results, err := models.Db.Query(
			"SELECT DISTINCT problems.id, title, problems.slug, difficulty, result FROM problems LEFT JOIN attempts ON problems.slug = attempts.slug AND username = $1 AND result = 'OK' WHERE difficulty = $2 ORDER BY title;",
			f.Username, f.Difficulty,
		)
		if err != nil {
			log.Println("failed to execute query", err)
		}
		for results.Next() {
			var p models.ProblemAndResult
			err = results.Scan(&p.Id, &p.Title, &p.Slug, &p.Difficulty, &result)
			if err != nil {
				log.Println("failed to scan", err)
			}

			if result.Valid {
				p.Result = result.String
			} else {
				p.Result = "-"
			}

			problems = append(problems, p)
		}

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		json.NewEncoder(w).Encode(problems)
	}

}

func ReturnProblem(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["slug"]

	var p models.Problem
	results, err := models.Db.Query("SELECT * FROM problems WHERE slug = $1;", keys)
	if err != nil {
		log.Println("failed to execute query", err)
	}
	for results.Next() {
		err = results.Scan(&p.Id, &p.Title, &p.Slug, &p.Difficulty, &p.Description, &p.Created)
		if err != nil {
			log.Println("failed to scan", err)
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(p)
}

func CheckProblem(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var input models.UserCode
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

	resp, err := http.Post("http://localhost:8091/run", "application/json", responseBody)
	if err != nil {
		fmt.Println("POST request failed: ", err)
		return
	}
	defer resp.Body.Close()

	var resFromContainer models.Result
	json.NewDecoder(resp.Body).Decode(&resFromContainer)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(resFromContainer)
}

func Submissions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reached submissions")
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var userSubmission models.Submission
	json.NewDecoder(r.Body).Decode(&userSubmission)
	fmt.Println(userSubmission.Username, userSubmission.Slug)

	var prevSubmissions = make([]models.Result, 0)
	results, err := models.Db.Query(
		"SELECT username, slug, lang, code, result, output FROM attempts WHERE username = $1 AND slug = $2;",
		userSubmission.Username, userSubmission.Slug,
	)
	if err != nil {
		log.Println("failed to query attempts", err)
	}
	for results.Next() {
		var prevSubmission models.Result
		err = results.Scan(
			&prevSubmission.Username, &prevSubmission.Slug, &prevSubmission.Lang,
			&prevSubmission.Code, &prevSubmission.Result, &prevSubmission.Output,
		)
		if err != nil {
			log.Println("failed to scan", err)
		}
		prevSubmissions = append(prevSubmissions, prevSubmission)
	}

	fmt.Println("hi", prevSubmissions)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(prevSubmissions)
}

func Solutions(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["slug"]

	var s models.Solution
	err := models.Db.QueryRow("SELECT solution FROM solutions WHERE slug = $1;", keys).Scan(&s.Solution)
	if err != nil {
		log.Println("failed to execute query", err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(s)
}

func Discussions(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["slug"]

	var discussions = make([]models.Discussion, 0)
	results, err := models.Db.Query("SELECT * FROM discussions WHERE slug = $1;", keys)
	if err != nil {
		log.Println("failed to execute query", err)
		return
	}
	for results.Next() {
		var d models.Discussion
		err = results.Scan(&d.Id, &d.Author, &d.Slug, &d.Title, &d.Description, &d.Created)
		if err != nil {
			log.Println("failed to scan", err)
		}
		discussions = append(discussions, d)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(discussions)
}

func Comments(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["discussionId"]

	var comments = make([]models.Comment, 0)
	results, err := models.Db.Query("SELECT * FROM comments WHERE discussion_id = $1;", keys)
	if err != nil {
		log.Println("failed to execute query, err")
	}
	for results.Next() {
		var c models.Comment
		err = results.Scan(&c.Id, &c.Author, &c.DiscussionId, &c.Description, &c.Created)
		if err != nil {
			log.Println("failed to scan", err)
		}
		comments = append(comments, c)
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(comments)
}

func NewComment(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["discussionId"]

	var newComment models.Comment
	json.NewDecoder(r.Body).Decode(&newComment)
	fmt.Println("new comment: ", newComment)

	err := models.Db.QueryRow(
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
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	keys := vars["username"]

	var u models.User
	err := models.Db.QueryRow(
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
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var newDiscussion models.Discussion
	json.NewDecoder(r.Body).Decode(&newDiscussion)
	fmt.Println("new discussion: ", newDiscussion)

	err := models.Db.QueryRow(
		"INSERT INTO discussions (author, slug, title, description) VALUES ($1, $2, $3, $4) RETURNING *;",
		&newDiscussion.Author, &newDiscussion.Slug, &newDiscussion.Title, &newDiscussion.Description,
	).Scan(&newDiscussion.Id, &newDiscussion.Author, &newDiscussion.Slug, &newDiscussion.Title, &newDiscussion.Description, &newDiscussion.Created)
	if err != nil {
		fmt.Println("failed to insert discussion: ", err)
		return
	}

	json.NewEncoder(w).Encode(&newDiscussion)
}
