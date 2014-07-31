package googletranslate

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var audioResponse = "bytesbytes"

func TestFetchSoundFile(t *testing.T) {
	server := buildTestServer(audioResponse)
	defer server.Close()

	err := fetchSoundFile(server.URL, "en", "lay down", "/tmp/gotr.test.mpg")
	assert.NoError(t, err)
	file, _ := ioutil.ReadFile("/tmp/gotr.test.mpg")

	assert.Equal(t, []byte(audioResponse), file)
}
