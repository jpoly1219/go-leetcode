package utils

import (
	"log"
	"os"
)

func SetLog() error {
	f, err := os.OpenFile("/backend/logs.txt", os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	return nil
}
