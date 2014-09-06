package cli

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var testLogRecord = LogRecord{
	From:          "en",
	To:            "pl",
	Phrase:        "turnstile",
	Translation:   "kołowrót",
	ExtraMeanings: "bramka, kołowrót przy wejściu",
}

func TestAddingToHisotory(t *testing.T) {
	historyFile, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}
	defer os.Remove(historyFile.Name())

	AddToHistory(historyFile.Name(), testLogRecord)
	AddToHistory(historyFile.Name(), testLogRecord)

	var historyObjects []LogRecord
	scanner := bufio.NewScanner(historyFile)

	for scanner.Scan() {
		var o LogRecord
		err := json.Unmarshal(scanner.Bytes(), &o)
		if err != nil {
			panic(err)
		}
		historyObjects = append(historyObjects, o)
	}

	if s := len(historyObjects); s != 2 {
		t.Errorf("Expected 2 log entries, got %v", s)
	}

	obj := historyObjects[0]

	if obj.From != testLogRecord.From {
		t.Errorf("Expected that From field is %v, got %v", testLogRecord.From, obj.From)
	}

	logTime, _ := time.Parse(time.RFC3339, obj.Date)
	nowPlus1Sec := time.Now().Local().Add(time.Second)
	logTimePlus2Sec := logTime.Add(2 * time.Second)

	if logTimePlus2Sec.Before(nowPlus1Sec) || logTime.After(nowPlus1Sec) {
		t.Errorf("Expected that Date field is properly stored, but got", obj.Date)
	}
}

func TestReadHistory(t *testing.T) {
	historyFile, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}
	defer os.Remove(historyFile.Name())

	serialisedJSON := "{'field':'fooo'}\n{'field':'bar'}"
	expectedJSON := "[{'field':'fooo'},{'field':'bar'}]"
	err = ioutil.WriteFile(historyFile.Name(), []byte(serialisedJSON), 0644)
	if err != nil {
		panic(err)
	}

	if actual, _ := ReadHistory(historyFile.Name()); actual != expectedJSON {
		t.Error("Expected history content to be equal", expectedJSON, ", but got:", actual)
	}
}

func TestReadHistoryWhenFileDoesntExists(t *testing.T) {
	_, err := ReadHistory("./foobarbizberdoesntexist")

	if err == nil {
		t.Error("Expected to return err, but got nothing")
	}
}
