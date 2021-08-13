package tests

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/jpoly1219/go-leetcode/pkg"
)

func TestRunCpp(t *testing.T) {
	type cppTest struct {
		input                  []byte
		dirUserfiles, expected string
	}

	pathUserfiles := filepath.Join(".", "testuserfiles")
	input1, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test1.cpp"))
	input2, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test2.cpp"))
	input3, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test3.cpp"))
	input4, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test4.cpp"))
	cppTests := []cppTest{
		{input1, pathUserfiles, "running test.cpp\n"},
		{input2, pathUserfiles, "foobaraboof\n"},
		{input3, pathUserfiles, "5 6 11 12 13 \n"},
		{input4, pathUserfiles, "The animal makes a sound \nThe pig says: wee wee \nThe dog says: bow wow \n"},
	}

	for _, test := range cppTests {
		output, _ := pkg.RunCpp(test.input, test.dirUserfiles)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestRunJava(t *testing.T) {
	type javaTest struct {
		input                  []byte
		dirUserfiles, expected string
	}

	pathUserfiles := filepath.Join(".", "testuserfiles")
	input1, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test1.java"))
	input2, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test2.java"))
	input3, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test3.java"))
	input4, _ := ioutil.ReadFile(filepath.Join(pathUserfiles, "test4.java"))
	javaTests := []javaTest{
		{input1, pathUserfiles, "running test.java\n"},
		{input2, pathUserfiles, "The string is a palindrome.\n"},
		{input3, pathUserfiles, "5 6 11 12 13 \n"},
		{input4, pathUserfiles, "The Shortest Superstring is gctaagttcatgcatc\n"},
	}

	for _, test := range javaTests {
		output, _ := pkg.RunJava(test.input, test.dirUserfiles)
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
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
