package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jpoly1219/go-leetcode/coderunner/lang"
	"github.com/jpoly1219/go-leetcode/coderunner/models"
	"github.com/jpoly1219/go-leetcode/coderunner/utils"
)

func RunTest(w http.ResponseWriter, r *http.Request) {
	// may be unnecessary when deploying the app using docker compose
	utils.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	type userCode struct {
		Username string `json:"username"`
		Slug     string `json:"slug"`
		Lang     string `json:"lang"`
		Code     string `json:"code"`
	}

	var code userCode
	json.NewDecoder(r.Body).Decode(&code)
	// fmt.Println("RunTest() reached: ", code.Username, code.Slug, code.Lang, code.Code)

	cwd, err := os.Getwd()
	fmt.Println(cwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	if filepath.Base(cwd) != "coderunner" {
		err = os.Chdir("..")
		if err != nil {
			fmt.Println("cd failed")
			fmt.Println(err)
			return
		}
	}

	var testcase string
	err = models.Db.QueryRow(
		"SELECT testcase FROM testcases WHERE slug = $1;",
		code.Slug,
	).Scan(&testcase)
	if err != nil {
		fmt.Println("failed to query and scan to testcase: ", err)
		return
	}

	testcasePath := "./testcase-" + code.Slug + ".json"
	err = os.WriteFile(testcasePath, []byte(testcase), 0644)
	if err != nil {
		fmt.Println("failed to create testcase.json")
	}

	var template string
	err = models.Db.QueryRow(
		"SELECT template FROM templates WHERE slug = $1 AND lang = $2;",
		code.Slug, code.Lang,
	).Scan(&template)
	if err != nil {
		fmt.Println("failed to query and scan to template: ", err)
		return
	}

	result, err := lang.HandleLangs(code.Username, code.Slug, code.Lang, code.Code, template)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*result)

	// save attempt to database
	_, err = models.Db.Exec(
		"INSERT INTO attempts (username, slug, lang, code, result, output) VALUES ($1, $2, $3, $4, $5, $6);",
		result.Username, result.Slug, result.Lang, result.Code, result.Result, result.Output,
	)
	if err != nil {
		fmt.Println("failed to insert attempt: ", err)
		return
	}
	json.NewEncoder(w).Encode(*result)
}
