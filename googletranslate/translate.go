package googletranslate

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const urlAddress = "https://translate.google.com/translate_a/t"

type Phrase struct {
	Translation   string   `json:"translation"`
	ExtraMeanings []string `json:"extra_meanings"`
}

func Translate(from string, to string, term string) (phrase Phrase, err error) {
	argError := func(code string) error {
		return fmt.Errorf("Unknown language code: %v. Check the list of available codes", code)
	}
	if inLangList(from) == false {
		err = argError(from)
		return
	}
	if inLangList(to) == false {
		err = argError(to)
		return
	}

	return OriginalTranslate(urlAddress, from, to, term)
}

func OriginalTranslate(urlAddress string, from string, to string, term string) (phrase Phrase, err error) {
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

	resp, err := runRquest(urlAddress, params)
	if err != nil {
		err = fmt.Errorf("Error fetching translation: [%v]", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	i1, i2 := getBracketIndexes(body)

	phrase.Translation = getTranslation(body, i1, i2)
	phrase.ExtraMeanings = getExtraMeanings(body, i1, i2)

	return
}

func getBracketIndexes(response []byte) (i1 []int, i2 []int) {
	var sep1 byte = 91 //[
	var sep2 byte = 93 //]

	for i, char := range response {
		if char == sep1 {
			i1 = append(i1, i)
		}
		if char == sep2 {
			i2 = append(i2, i)
		}
	}
	return
}

func getExtraMeanings(response []byte, i1 []int, i2 []int) (translations []string) {
	pos1 := i1[5] + 1
	pos2 := i2[2]
	tempStr := strings.Replace(string(response[pos1:pos2]), "\"", "", -1)

	// this is a guard for a situation when someone wants to translate a sentence and there are no extra meanings
	if len(tempStr) < 2 {
		return
	}

	translations = strings.Split(tempStr, ",")
	return
}

func getTranslation(response []byte, i1 []int, i2 []int) string {
	pos1 := i1[2] + 1
	pos2 := i2[0]

	arry := strings.Split(string(response[pos1:pos2]), "\",\"")
	return strings.Replace(arry[0], "\"", "", -1)
}
