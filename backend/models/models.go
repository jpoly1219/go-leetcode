package models

import (
	"database/sql"

	"github.com/google/uuid"
)

var Db *sql.DB

type Problem struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Difficulty  string    `json:"difficulty"`
	Description string    `json:"description"`
	Created     string    `json:"created"`
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
	Id         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Difficulty string    `json:"difficulty"`
	Result     string    `json:"result"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	Userid     uuid.UUID `json:"userid"`
	Username   string    `json:"username"`
	Fullname   string    `json:"fullname"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	ProfilePic string    `json:"profilePic"`
}

type NewUserData struct {
	OldUsername   string `json:"oldUsername"`
	NewUsername   string `json:"newUsername"`
	NewFullname   string `json:"newFullname"`
	NewEmail      string `json:"newEmail"`
	NewPassword   string `json:"newPassword"`
	NewProfilePic string `json:"newProfilePic"`
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
	Id          uuid.UUID `json:"id"`
	Author      string    `json:"author"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     string    `json:"created"`
}

type Comment struct {
	Id           uuid.UUID `json:"id"`
	Author       string    `json:"author"`
	DiscussionId string    `json:"discussionId"`
	Description  string    `json:"description"`
	Created      string    `json:"created"`
}

type Filter struct {
	Username   string `json:"username"`
	Difficulty string `json:"difficulty"`
}
