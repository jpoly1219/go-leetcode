package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func HandleCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://44.207.2.100:8090")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func FileToLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	return LinesFromFile(f)
}

func LinesFromFile(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func WriteCodeToFile(filePath, code string, lines []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	codeLines := strings.Split(code, "\n")

	writer := bufio.NewWriter(f)
	for _, line := range lines {
		if strings.Contains(line, "insert Solution class here") {
			for _, codeLine := range codeLines {
				_, _ = writer.WriteString(codeLine + "\n")
			}
		}
		_, _ = writer.WriteString(line + "\n")
	}
	writer.Flush()
	return nil
}
