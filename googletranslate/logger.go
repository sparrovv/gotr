package googletranslate

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "gotr: ", log.Ldate|log.Ltime)
}

func Debug(txt string) {
	if os.Getenv("GOTR_DEBUG") == "y" {
		logger.Println(txt)
	}
}
