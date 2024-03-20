package cli

import (
	"fmt"
	"log"
	"os"
)

func GetLogger(name string, filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logfile, fmt.Sprintf("(%s) ", name), log.Ldate|log.Ltime)
}
