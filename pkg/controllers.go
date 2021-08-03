package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint reached")

	baseDir := LoadEnv("BASEDIR")
	userfilesDir := filepath.Join(baseDir, "userfiles/")
	content, _ := ioutil.ReadFile(filepath.Join(userfilesDir, "test.cpp"))
	runCpp(content, userfilesDir)
	fmt.Fprintf(w, "done\n")

	/*
		out1 := make(chan string)
		out2 := make(chan string)

		go PrintHello(5, out1)
		go PrintBye(5, out2)

		for i := 0; i < 10; i++ {
			select {
			case msg := <-out1:
				fmt.Print(msg)
				fmt.Fprint(w, msg)
			case msg := <-out2:
				fmt.Print(msg)
				fmt.Fprint(w, msg)
			}
		}

		fmt.Fprintf(w, "done\n")
	*/
}
