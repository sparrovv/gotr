package cli

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type LogRecord struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Phrase        string `json:"phrase"`
	Translation   string `json:"translation"`
	ExtraMeanings string `json:"extraMeanings"`
	Date          string `json:"Date"`
}

func AddToHistory(logPath string, lr LogRecord) {
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	l := log.New(f, "", 0)

	lr.Date = time.Now().Local().Format(time.RFC3339)

	b, err := json.Marshal(lr)
	if err != nil {
		fmt.Println(err)
	}

	l.Println(string(b))
}

func ReadHistory(logPath string) (history string, err error) {
	f, err := os.Open(logPath)
	if err != nil {
		return history, errors.New("History file doesn't exists")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		history += scanner.Text() + ","
	}

	history = "[" + history[:len(history)-1] + "]"

	return
}
