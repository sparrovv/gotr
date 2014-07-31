package googletranslate

import (
	"net/http"
	"net/url"
)

func runRquest(urlString string, params map[string]string) (resp *http.Response, err error) {
	urlObj, err := url.Parse(urlString)
	check(err)

	v := url.Values{}

	for key, value := range params {
		v.Set(key, value)
	}

	urlObj.RawQuery = v.Encode()

	req, err := http.NewRequest("GET", urlObj.String(), nil)
	check(err)

	req.Header.Add("User-Agent", "Mozilla/5.0")

	return http.DefaultClient.Do(req)
}
