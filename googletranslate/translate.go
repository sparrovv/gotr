package googletranslate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Phrase struct {
	Translation   string   `json:"translation"`
	ExtraMeanings []string `json:"extra_meanings"`
}

func (p Phrase) Result() string {
	return fmt.Sprintf("translation:\n%s - %v", p.Translation, strings.Join(p.ExtraMeanings, ", "))
}

func (p Phrase) JsonResult() string {
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func Translate(urlAddress string, from string, to string, term string) (phrase Phrase, err error) {
	qparams := map[string]string{
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

	urlObj, err := url.Parse(urlAddress)
	if err != nil {
		log.Fatal(err)
	}
	v := url.Values{}

	for key, value := range qparams {
		v.Set(key, value)
	}

	urlObj.RawQuery = v.Encode()

	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		return
	}

	req.Header.Add("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = fmt.Errorf("Error fetching assignments: [%v]", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

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

	// this is a guard for a situation when someone want's to translate a sentence and there are no extra meanings
	if len(tempStr) < 2 {
		return
	}

	translations = strings.Split(tempStr, ",")
	return
}

func getTranslation(response []byte, i1 []int, i2 []int) string {
	pos1 := i1[2] + 1
	pos2 := i2[0]

	arry := strings.Split(strings.Replace(string(response[pos1:pos2]), "\"", "", -1), ",")
	return arry[0]
}
