package pkg

import (
	"log"
	"os"
)

func writeLog(message string, err error) {
	f, _ := os.OpenFile("/backend/logs.txt", os.O_APPEND, 0644)
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println(message, err)
}
