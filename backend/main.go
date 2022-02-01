package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jpoly1219/go-leetcode/backend/auth"
	"github.com/jpoly1219/go-leetcode/backend/controllers"
	"github.com/jpoly1219/go-leetcode/backend/middlewares"
	"github.com/jpoly1219/go-leetcode/backend/models"
	"github.com/jpoly1219/go-leetcode/backend/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := utils.SetLog()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("app start")
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file.")
	}

	const (
		host     = "database"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "goleetcode"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	models.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	defer models.Db.Close()

	r := mux.NewRouter()
	// r.HandleFunc("/problemsets", controllers.Problemsets)
	r.HandleFunc("/submissions", controllers.Submissions)

	problemsetsR := r.PathPrefix("/problemsets").Subrouter()
	problemsetsR.Handle("/all", middlewares.VerifyToken(http.HandlerFunc(controllers.ProblemsetsAll)))
	problemsetsR.Handle("/filter", middlewares.VerifyToken(http.HandlerFunc(controllers.ProblemsetsFilter)))

	solveR := r.PathPrefix("/solve").Subrouter()
	solveR.Handle("/{slug}", middlewares.VerifyToken(http.HandlerFunc(controllers.SolveSlug)))

	solutionsR := r.PathPrefix("/solutions").Subrouter()
	solutionsR.HandleFunc("/{slug}", controllers.SolutionsSlug)

	discussionsR := r.PathPrefix("/discussions").Subrouter()
	discussionsR.HandleFunc("/newdiscussion", controllers.DiscussionsNewdiscussion)
	discussionsR.HandleFunc("/{slug}", controllers.DiscussionsSlug)
	discussionsR.HandleFunc("/{slug}/{discussionId}", controllers.DiscussionsSlugDiscussionid)
	discussionsR.HandleFunc("/{slug}/{discussionId}/newcomment", controllers.DiscussionsSlugDiscussionidNewcomment)

	checkR := r.PathPrefix("/check").Subrouter()
	checkR.HandleFunc("/{slug}", controllers.CheckSlug)

	authR := r.PathPrefix("/auth").Subrouter()
	authR.HandleFunc("/signup", auth.Signup)
	authR.HandleFunc("/login", auth.Login)
	authR.HandleFunc("/silentrefresh", auth.SilentRefresh)

	usersR := r.PathPrefix("/users").Subrouter()
	usersR.HandleFunc("/{username}", controllers.UsersUsername)

	log.Fatal(http.ListenAndServe(":8090", r))
}
