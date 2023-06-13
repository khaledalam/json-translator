package jsontranslator

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

// Skip test case if no internet connection
func checkInternetConnection(t *testing.T) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		log.Println("Skipped: no internet connection.")
		t.SkipNow()
	}
}

func TestTranslateJsonFile(t *testing.T) {
	checkInternetConnection(t)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	jsonFile := path + "/cli/test.json"
	languageFrom := "en"
	languageTo := "de"

	result := TranslateJsonFile(jsonFile, 5, languageFrom, languageTo)

	expected := `{
  "Account": "Konto",
  "Home": "Heim"
}`
	assert.Equal(t, expected, string(result), "Must be equal.")
	log.Println("-----")
}

func TestTranslateJsonFiles(t *testing.T) {

	checkInternetConnection(t)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	jsonFile := path + "/cli/test.json"
	languageFrom := "en"
	languagesTo := []string{"de", "it"}

	result := TranslateJsonFiles(jsonFile, 5, languageFrom, languagesTo...)

	expected := `{
  "de": {
    "Account": "Konto",
    "Home": "Heim"
  },
  "it": {
    "Account": "Account",
    "Home": "Casa"
  }
}`
	assert.Equal(t, expected, string(result), "Must be equal.")
	log.Println("-----")
}

func TestInvalidInputJsonFile(t *testing.T) {
	checkInternetConnection(t)

	languageFrom := "en"
	languagesTo := []string{"de", "it"}

	result := TranslateJsonFiles("invalid/json/file/path/a.json", 5, languageFrom, languagesTo...)

	assert.Equal(t, []byte(nil), result, "Must be equal.")
	log.Println("-----")
}
