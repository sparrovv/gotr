package googletranslate

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const speechURL = "https://translate.google.com/translate_tts"

func FetchSoundFile(lang string, text string, audioPath string) (err error) {
	return fetchSoundFile(speechURL, lang, text, audioPath)
}

func fetchSoundFile(url string, lang string, text string, audioPath string) (err error) {
	params := map[string]string{
		"text": text,
		"tl":   lang,
		"ie":   "UTF-8",
		"oe":   "UTF-8",
	}

	multipleParams := map[string][]string{} // emtpy

	resp, err := runRquest(url, params, multipleParams)
	if err != nil {
		err = fmt.Errorf("Error fetching translation: [%v]", err)
		return
	}
	if resp.StatusCode == http.StatusNotFound {
		err = errors.New("Speech synthesis is not supported for this language: " + lang)
		return
	}

	if resp.StatusCode == http.StatusServiceUnavailable {
		err = errors.New("Google has detetect an invalid request... Sorry :/")
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	defer resp.Body.Close()

	err = ioutil.WriteFile(audioPath, body, 0644)
	check(err)

	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
