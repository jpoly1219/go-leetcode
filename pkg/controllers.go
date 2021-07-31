package controllers

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint reached")
	w.Write([]byte("Hello world!\n"))
}
