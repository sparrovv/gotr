package googletranslate

import (
	"net/http"
	"net/url"
	"time"
)

// Default client timetout to all google requests
var clientTimeout = time.Duration(10 * time.Second)

func runRquest(urlString string, standardParams map[string]string, multipleValuesParams map[string][]string) (resp *http.Response, err error) {
	urlObj, err := url.Parse(urlString)
	check(err)

	queryString := url.Values{}

	for key, value := range standardParams {
		queryString.Set(key, value)
	}

	for key, values := range multipleValuesParams {
		for _, value := range values {
			queryString.Add(key, value)
		}
	}

	urlObj.RawQuery = queryString.Encode()

	Debug("request: " + urlObj.String())
	req, err := http.NewRequest("POST", urlObj.String(), nil)
	check(err)

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; OpenBSD i386) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.125 Safari/537.36")

	client := http.Client{
		Timeout: clientTimeout,
	}

	return client.Do(req)
}
