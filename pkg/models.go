package pkg

import (
	"database/sql"
)

var Db *sql.DB

type problem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Difficulty  string `json:"difficulty"`
	Description string `json:"description"`
	Created     string `json:"created"`
}

type userCode struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
}
