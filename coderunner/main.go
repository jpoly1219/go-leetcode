package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jpoly1219/go-leetcode/coderunner/controllers"
	"github.com/jpoly1219/go-leetcode/coderunner/models"
	_ "github.com/lib/pq"
)

func main() {
	// establish database connection
	var err error
	const (
		host     = "go-leetcode-rds.cvbgqae2t74s.us-east-1.rds.amazonaws.com"
		port     = 5432
		user     = "postgres"
		password = "NT9oOSqLJWcNUIRuJv35"
		dbname   = "goleetcode"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	models.Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect to db")
	}
	defer models.Db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/coderunner/run", controllers.RunTest)

	// each container will have the exact same domain, so how will the backend distinguish between the containers?
	// plus, if all containers use port 8091, it would be even more confusing...
	// may need a container orchestration tool
	log.Fatal(http.ListenAndServe(":8091", r))
}
