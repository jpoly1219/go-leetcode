package pkg

import (
	"database/sql"
)

var Db *sql.DB

type problem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Difficulty  string `json:"difficulty"`
	Description string `json:"description"`
}
