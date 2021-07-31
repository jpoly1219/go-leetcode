package pkg

import (
	"time"
)

func PrintHello(num int, out chan string) {
	defer close(out)
	for i := 0; i < num; i++ {
		out <- "hello world!\n"
		time.Sleep(time.Second)
	}
}

func PrintBye(num int, out chan string) {
	defer close(out)
	for i := 0; i < num; i++ {
		out <- "goodbye!\n"
		time.Sleep(time.Second)
	}
}
