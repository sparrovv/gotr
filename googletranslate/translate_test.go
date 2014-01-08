package googletranslate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var googleTranslateResponse = `
[[["położyć","lay down","",""]],[["verb",["położyć","kłaść","pokłaść","układać","składać","postawić","zakładać","narysować"],[["położyć",["put","place","down","lay down","set","put down"],,0.002638618],["kłaść",["place","put","lay","lay down","set","build"],,0.00108288659],["pokłaść",["lay down"],,0.000458523049],["układać",["lay","place","compose","recline","lay down","negotiate"],,0.000273796613],["składać",["fold","assemble","compose","lay down","render","consign"],,0.000188177481],["postawić",["put","place","set","raise","erect","lay down"],,0.00011593278],["zakładać",["found","establish","set up","start","imply","lay down"],,6.20543433e-05],["narysować",["draw","trace","draw out","lay down","portray","crayon"],,2.93123976e-05]],"lay down",2]],"en",,[["położyć",[4],1,0,1000,0,1,0]],[["lay down",4,[["położyć",1000,1,0],["ustanawiają",0,1,0],["ustanowić",0,1,0],["ustanowienie",0,1,0],["określają",0,1,0]],[[0,8]],"lay down"]],,[,"lie down",[10],,1],[],68]
`

var fetchHandler = func(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(rw, googleTranslateResponse)
}

func TestTranslate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(fetchHandler))

	phrase, err := Translate(server.URL, "en", "pl", "turn up")

	assert.NoError(t, err)
	assert.Equal(t, phrase.ExtraMeanings, []string{"położyć", "kłaść", "pokłaść", "układać", "składać", "postawić", "zakładać", "narysować"})
	assert.Equal(t, phrase.Translation, "położyć")

	server.Close()
}

func TestResult(t *testing.T) {
	phrase := Phrase{Translation: "Foo", ExtraMeanings: []string{"foo", "bar"}}

	assert.Equal(t, phrase.Result(), "translation:\nFoo - foo, bar")
}
