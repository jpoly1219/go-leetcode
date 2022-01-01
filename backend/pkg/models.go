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
	Slug     string `json:"slug"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
	Result   string `json:"result"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Output   string `json:"output"`
}

type problemAndResult struct {
	Id         int    `json:"id"`
	Title      string `json:"string"`
	Slug       string `json:"slug"`
	Difficulty string `json:"difficulty"`
	Result     string `json:"result"`
}

type token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type user struct {
	Userid     int    `json:"userid"`
	Username   string `json:"username"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfilePic string `json:"profilePic"`
}

type submission struct {
	Username string `json:"username"`
	Slug     string `json:"slug"`
}

type solution struct {
	Slug     string `json:"slug"`
	Solution string `json:"solution"`
}

type discussion struct {
	Id          int    `json:"id"`
	Author      string `json:"author"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
}

type comment struct {
	Id           int    `json:"id"`
	Author       string `json:"author"`
	DiscussionId string `json:"discussionId"`
	Description  string `json:"description"`
	Created      string `json:"created"`
}

type filter struct {
	Difficulty string `json:"difficulty"`
}
