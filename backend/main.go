package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jpoly1219/go-leetcode/backend/controllers"
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
		host     = "jpoly1219devbox.xyz"
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
	r.HandleFunc("/problemsets", controllers.Problemsets)
	r.HandleFunc("/submissions", controllers.Submissions)

	problemsetsR := r.PathPrefix("/problemsets").Subrouter()
	problemsetsR.HandleFunc("/all", controllers.Problemsets)
	problemsetsR.HandleFunc("/filter", controllers.FilterProblemsets)

	solveR := r.PathPrefix("/solve").Subrouter()
	solveR.Handle("/{slug}", controllers.VerifyToken(http.HandlerFunc(controllers.ReturnProblem)))

	solutionsR := r.PathPrefix("/solutions").Subrouter()
	solutionsR.HandleFunc("/{slug}", controllers.Solutions)

	discussionsR := r.PathPrefix("/discussions").Subrouter()
	discussionsR.HandleFunc("/newdiscussion", controllers.NewDiscussion)
	discussionsR.HandleFunc("/{slug}", controllers.Discussions)
	discussionsR.HandleFunc("/{slug}/{discussionId}", controllers.Comments)
	discussionsR.HandleFunc("/{slug}/{discussionId}/newcomment", controllers.NewComment)

	checkR := r.PathPrefix("/check").Subrouter()
	checkR.HandleFunc("/{slug}", controllers.CheckProblem)

	authR := r.PathPrefix("/auth").Subrouter()
	authR.HandleFunc("/signup", controllers.Signup)
	authR.HandleFunc("/login", controllers.Login)
	authR.HandleFunc("/silentrefresh", controllers.SilentRefresh)

	usersR := r.PathPrefix("/users").Subrouter()
	usersR.HandleFunc("/{username}", controllers.GetUser)

	log.Fatal(http.ListenAndServe(":8090", r))
}
