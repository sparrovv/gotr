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

var googleTranslateResponseWithSentence = `
[[["Therefore, it is worth testing , because otherwise","Dlatego też warto testować, bo inaczej","",""]],,"pl",,[["Therefore,",[4],true,false,940,0,2,0],["it is worth",[5],true,false,940,2,5,0],["testing",[6],true,false,947,5,6,0],[", because otherwise",[7],false,false,512,6,9,0]],[["Dlatego też",4,[["Therefore,",940,true,false],["Therefore",23,true,false],["That is why",18,true,false],["It is therefore",9,true,false],["Thus",0,true,false]],[[0,11]],"Dlatego też warto testować, bo inaczej"],["warto",5,[["it is worth",940,true,false],["worth",9,true,false],["value",0,true,false],["value of",0,true,false],["is worth",0,true,false]],[[12,17]],""],["testować",6,[["testing",947,true,false],["test",0,true,false],["tested",0,true,false],["to test",0,true,false],["test the",0,true,false]],[[18,26]],""],[", bo inaczej",7,[[", because otherwise",512,false,false],["or else",0,false,false]],[[26,38]],""]],,,[["pl"]],33]
`

func buildTestServer(response string) *httptest.Server {
	var fetchHandler = func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintf(rw, response)
	}

	return httptest.NewServer(http.HandlerFunc(fetchHandler))
}

func TestTranslateWithSimpleResponse(t *testing.T) {
	server := buildTestServer(googleTranslateResponse)
	defer server.Close()

	phrase, err := OriginalTranslate(server.URL, "en", "pl", "lay down")

	assert.NoError(t, err)
	assert.Equal(t, phrase.ExtraMeanings, []string{"położyć", "kłaść", "pokłaść", "układać", "składać", "postawić", "zakładać", "narysować"})
	assert.Equal(t, phrase.Translation, "położyć")
}

func TestTranslateWithFullSentenceResponse(t *testing.T) {
	server := buildTestServer(googleTranslateResponseWithSentence)
	defer server.Close()

	phrase, err := OriginalTranslate(server.URL, "pl", "en", "Dlatego też warto testować, bo inaczej")

	assert.NoError(t, err)
	assert.Equal(t, phrase.Translation, "Therefore, it is worth testing , because otherwise")
	assert.Equal(t, len(phrase.ExtraMeanings), 0)
}

func TestTranslateWithWrongLangCode(t *testing.T) {
	_, err := Translate("ende", "pl", "lay down")

	assert.NotNil(t, err)
	assert.Equal(t, "Unknown language code: ende. Check the list of available codes", err.Error())
}
