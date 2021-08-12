package tests

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/jpoly1219/go-leetcode/pkg"
)

func TestRunCpp(t *testing.T) {
	pathUserfiles := filepath.Join(".", "testuserfiles")
	input, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test.cpp"))
	output, _ := pkg.RunCpp(input, pathUserfiles)
	if output != "running test.cpp\n" {
		t.Errorf("expected 'running test.cpp'")
	}
}

func TestRunJava(t *testing.T) {
	pathUserfiles := filepath.Join(".", "testuserfiles")
	input, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test.java"))
	output, _ := pkg.RunJava(input, pathUserfiles)
	if output != "running test.java\n" {
		t.Errorf("expected 'running test.java'")
	}
}

func TestRunJs(t *testing.T) {
	pathUserfiles := filepath.Join(".", "testuserfiles")
	input, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test.js"))
	output, _ := pkg.RunJs(input, pathUserfiles)
	if output != "running test.js\n" {
		t.Errorf("expected 'running test.js'")
	}
}

func TestRunPy(t *testing.T) {
	pathUserfiles := filepath.Join(".", "testuserfiles")
	input, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test.py"))
	output, _ := pkg.RunPy(input, pathUserfiles)
	if output != "running test.py\n" {
		t.Errorf("expected 'running test.py'")
	}
}
