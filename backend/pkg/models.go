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
	Username string `json:"username"`
	Slug     string `json:"slug"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
}

type result struct {
	Username string `json:"username"`
	Slug     string `json:"Slug"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
	Result   string `json:"result"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Output   string `json:"output"`
}

type token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type user struct {
	Userid   int    `json:"userid"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type submission struct {
	Username string `json:"username"`
	Slug     string `json:"slug"`
}
