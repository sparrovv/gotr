package googletranslate

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const translateURL = "https://translate.google.com/translate_a/t"

type Phrase struct {
	Translation   string   `json:"translation"`
	ExtraMeanings []string `json:"extra_meanings"`
}

func Translate(from string, to string, term string) (phrase Phrase, err error) {
	argError := func(code string) error {
		return fmt.Errorf("Unknown language code: %v. Check the list of available codes", code)
	}
	if isInLangList(from) == false {
		err = argError(from)
		return
	}
	if isInLangList(to) == false {
		err = argError(to)
		return
	}

	return translate(translateURL, from, to, term)
}

func translate(translateURL string, from string, to string, term string) (phrase Phrase, err error) {
	params := map[string]string{
		"client":   "t",
		"hl":       "en",
		"multires": "1",
		"sc":       "1",
		"sl":       from,
		"ssel":     "0",
		"tl":       to,
		"tsel":     "0",
		"uptl":     "en",
		"text":     term,
	}

	resp, err := runRquest(translateURL, params)
	if err != nil {
		err = fmt.Errorf("Error fetching translation: [%v]", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	phrase.Translation = getTranslation(string(body))
	phrase.ExtraMeanings = getExtraMeanings(string(body))

	return
}

func getTranslation(response string) string {
	dirtyTranslation := strings.Split(response, "[[")[1]
	return strings.Split(dirtyTranslation, "\"")[1]
}

func getExtraMeanings(response string) (extraMeanings []string) {
	tempSplit := strings.Split(strings.Split(response, "[[")[2], "[")

	if len(tempSplit) <= 1 {
		return
	}

	possibleSynonyms := tempSplit[1]
	length := len(possibleSynonyms)

	if !strings.Contains(possibleSynonyms, "true,false") && !strings.Contains(possibleSynonyms, "false,false") {
		sStrings := strings.Replace(possibleSynonyms[:length-3], "\"", "", -1)

		extraMeanings = strings.Split(sStrings, ",")
	}

	return
}
