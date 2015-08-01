package googletranslate

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var audioResponse = "bytesbytes"

func TestFetchSoundFile(t *testing.T) {
	server := buildTestServer(audioResponse, http.StatusOK)
	defer server.Close()

	err := fetchSoundFile(server.URL, "en", "lay down", "/tmp/gotr.test.mpg")
	assert.NoError(t, err)
	file, _ := ioutil.ReadFile("/tmp/gotr.test.mpg")

	assert.Equal(t, []byte(audioResponse), file)
}

func TestFetchSoundFileWhen404(t *testing.T) {
	server := buildTestServer(audioResponse, http.StatusNotFound)
	defer server.Close()

	err := fetchSoundFile(server.URL, "en", "lay down", "/tmp/gotr.test.mpg")
	assert.Equal(t, err.Error(), "Speech synthesis is not supported for this language: en")
}

func TestFetchSoundFileWhen503(t *testing.T) {
	server := buildTestServer(audioResponse, http.StatusServiceUnavailable)
	defer server.Close()

	err := fetchSoundFile(server.URL, "en", "lay down", "/tmp/gotr.test.mpg")
	assert.Equal(t, err.Error(), "Google has detetect an invalid request... Sorry :/")
}
