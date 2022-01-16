package models

import (
	"database/sql"
)

var Db *sql.DB

type Problem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Difficulty  string `json:"difficulty"`
	Description string `json:"description"`
	Created     string `json:"created"`
}

type UserCode struct {
	Username string `json:"username"`
	Slug     string `json:"slug"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
}

type Result struct {
	Username string `json:"username"`
	Slug     string `json:"slug"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
	Result   string `json:"result"`
	Input    string `json:"input"`
	Expected string `json:"expected"`
	Output   string `json:"output"`
}

type ProblemAndResult struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Difficulty string `json:"difficulty"`
	Result     string `json:"result"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	Userid     int    `json:"userid"`
	Username   string `json:"username"`
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ProfilePic string `json:"profilePic"`
}

type Submission struct {
	Username string `json:"username"`
	Slug     string `json:"slug"`
}

type Solution struct {
	Slug     string `json:"slug"`
	Solution string `json:"solution"`
}

type Discussion struct {
	Id          int    `json:"id"`
	Author      string `json:"author"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     string `json:"created"`
}

type Comment struct {
	Id           int    `json:"id"`
	Author       string `json:"author"`
	DiscussionId string `json:"discussionId"`
	Description  string `json:"description"`
	Created      string `json:"created"`
}

type Filter struct {
	Username   string `json:"username"`
	Difficulty string `json:"difficulty"`
}
