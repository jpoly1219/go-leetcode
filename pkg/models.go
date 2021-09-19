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

type result struct {
	Username string `json:"username"`
	Pnum     int    `json:"pnum"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
	Result   string `json:"result"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Output   string `json:"output"`
	// Runtime string `json:"runtime"`
}

type token struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type user struct {
	Userid string `json:"userid"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
}