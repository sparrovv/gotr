package googletranslate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
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

	sanitizedBody := sanitizeBody(body)

	var translateResponse []interface{}
	err = json.Unmarshal(sanitizedBody, &translateResponse)
	if err != nil {
		log.Fatal(err)
	}

	phrase.Translation = getTranslation(translateResponse)
	phrase.ExtraMeanings = getExtraMeanings(translateResponse)

	return
}

// Get rid of invalid chars ",," or "[,"
// Google translate response is not a JSON, but looks almost as it was.
func sanitizeBody(body []byte) (sanitizedBody []byte) {
	reg1, err := regexp.Compile(",{2,}")
	if err != nil {
		log.Fatal(err)
	}
	reg2, err := regexp.Compile(`\[,{1,}`)
	if err != nil {
		log.Fatal(err)
	}

	sanitizedBody = reg1.ReplaceAll(body, []byte(","))
	sanitizedBody = reg2.ReplaceAll(sanitizedBody, []byte(`[`))
	return
}

func getExtraMeanings(foo []interface{}) (translations []string) {
	secondArray := foo[1]
	insideArray, ok := secondArray.([]interface{})
	if ok {
		fizArray := insideArray[0]
		insideArray, ok := fizArray.([]interface{})
		if ok {
			bozArray := insideArray[1]
			stringArray, ok := bozArray.([]interface{})
			if ok {
				for _, str := range stringArray {
					translations = append(translations, str.(string))
				}
			}
		}
	}

	return
}

func getTranslation(foo []interface{}) (translation string) {
	firstArray := foo[0]
	insideArray, ok := firstArray.([]interface{})
	if ok {
		fiz := insideArray[0]
		stringArray, ok := fiz.([]interface{})
		if ok {
			translation = stringArray[0].(string)
		}
	}

	return
}
