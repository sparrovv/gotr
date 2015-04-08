package googletranslate

import (
	"net/http"
	"net/url"
	"time"
)

// Default client timetout value to all google requests
var clientTimeout = time.Duration(10 * time.Second)

func runRquest(urlString string, params map[string]string) (resp *http.Response, err error) {
	urlObj, err := url.Parse(urlString)
	check(err)

	v := url.Values{}

	for key, value := range params {
		v.Set(key, value)
	}

	urlObj.RawQuery = v.Encode()

	req, err := http.NewRequest("POST", urlObj.String(), nil)
	check(err)

	req.Header.Add("User-Agent", "Mozilla/5.0")

	client := http.Client{
		Timeout: clientTimeout,
	}

	return client.Do(req)
}
