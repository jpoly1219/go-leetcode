package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint reached")

	userfilesDir := filepath.Join(".", "userfiles/")
	content, _ := ioutil.ReadFile(filepath.Join(userfilesDir, "test.cpp"))
	outCpp, err := runCpp(content, userfilesDir)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(outCpp))

}
