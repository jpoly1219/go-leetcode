package models

import "database/sql"

var Db *sql.DB

type ResultFile struct {
	Username string `json:"username"`
	Slug     string `json:"Slug"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
	Result   string `json:"result"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Output   string `json:"output"`
	// Runtime string `json:"runtime"`
}
