package pkg

import (
	"log"
	"os"
)

func writeLog(message string, err error) error {
	f, err := os.OpenFile("/", os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println(message, err)

	return nil
}
