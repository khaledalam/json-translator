package jsontranslator

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestTranslateJsonFile(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	jsonFile := path + "/cli/test.json"
	languageFrom := "en"
	languageTo := "de"

	result := TranslateJsonFile(jsonFile, 5, languageFrom, languageTo)

	must := `{
  "Account": "Konto",
  "Home": "Heim"
}`
	assert.Equal(t, must, string(result), "Must be equal.")
	log.Println("-----")
}

func TestTranslateJsonFiles(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	jsonFile := path + "/cli/test.json"
	languageFrom := "en"
	languagesTo := []string{"de", "it"}

	result := TranslateJsonFiles(jsonFile, 5, languageFrom, languagesTo...)

	must := `{
  "de": {
    "Account": "Konto",
    "Home": "Heim"
  },
  "it": {
    "Account": "Account",
    "Home": "Casa"
  }
}`
	assert.Equal(t, must, string(result), "Must be equal.")
	log.Println("-----")

}
