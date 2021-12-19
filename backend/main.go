package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/jpoly1219/go-leetcode/backend/pkg"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("app start")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file.")
	}

	const (
		host     = "jpoly1219devbox.xyz"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "goleetcode"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	pkg.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	defer pkg.Db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/run", pkg.Run)
	r.HandleFunc("/problemsets", pkg.Problemsets)
	r.HandleFunc("/submissions", pkg.Submissions)

	problemsetsR := r.PathPrefix("/problemsets").Subrouter()
	problemsetsR.HandleFunc("/filter", pkg.FilterProblemsets)

	solveR := r.PathPrefix("/solve").Subrouter()
	solveR.Handle("/{slug}", pkg.VerifyToken(http.HandlerFunc(pkg.ReturnProblem)))
	// solveR.HandleFunc("/{slug}", pkg.ReturnProblem)

	solutionsR := r.PathPrefix("/solutions").Subrouter()
	solutionsR.HandleFunc("/{slug}", pkg.Solutions)

	discussionsR := r.PathPrefix("/discussions").Subrouter()
	discussionsR.HandleFunc("/newdiscussion", pkg.NewDiscussion)
	discussionsR.HandleFunc("/{slug}", pkg.Discussions)
	discussionsR.HandleFunc("/{slug}/{discussionId}", pkg.Comments)
	discussionsR.HandleFunc("/{slug}/{discussionId}/newcomment", pkg.NewComment)

	checkR := r.PathPrefix("/check").Subrouter()
	checkR.HandleFunc("/{slug}", pkg.CheckProblem)

	authR := r.PathPrefix("/auth").Subrouter()
	authR.HandleFunc("/signup", pkg.Signup)
	authR.HandleFunc("/login", pkg.Login)
	authR.HandleFunc("/silentrefresh", pkg.SilentRefresh)

	usersR := r.PathPrefix("/users").Subrouter()
	usersR.HandleFunc("/{username}", pkg.GetUser)

	log.Fatal(http.ListenAndServe(":8090", r))
}
