package googletranslate

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// google translate URL
const translateURL = "https://translate.google.com/translate_a/single"

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
	standardParams := map[string]string{
		"client": "t",
		"ie":     "UTF-8",
		"oe":     "UTF-8",
		"sl":     from,
		"tl":     to,
		"q":      term,
		"pc":     "1",
		"otf":    "1",
		"srcrom": "1",
		"ssel":   "0",
		"tsel":   "0",
	}

	multipleValuesParams := map[string][]string{
		"dt": []string{"bd", "ex", "ld", "md", "qc", "rw", "rm", "ss", "t", "at"},
	}

	resp, err := runRquest(translateURL, standardParams, multipleValuesParams)
	if err != nil {
		err = fmt.Errorf("Error fetching translation: [%v]", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	phrase.Translation = getTranslation(string(body))
	// Get extra meanings only when "from" language is different than "to" language
	if from != to {
		phrase.ExtraMeanings = getExtraMeanings(string(body))
	}
	return
}

func getTranslation(response string) string {
	dirtyTranslation := strings.Split(response, "[[")[1]
	return strings.Split(dirtyTranslation, "\"")[1]
}

func getExtraMeanings(response string) (extraMeanings []string) {
	dblSquareBracketSplit := strings.Split(response, "[[")
	if len(dblSquareBracketSplit) < 3 {
		return
	}

	tempSplit := strings.Split(dblSquareBracketSplit[2], "[")
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
