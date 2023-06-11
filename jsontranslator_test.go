package jsontranslator

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestTranslateJson(t *testing.T) {

	jsonFile := "test.json"
	langFrom := "en"
	langTo := "de"

	// Let's first read the `config.json` file
	content, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal("[Testing]: Error when opening json file: ", err)
	}

	result := TranslateJson(content, langFrom, langTo)

	must := `{
  "Account": "Konto",
  "Home": "Heim"
}`
	assert.Equal(t, must, string(result), "Must be equal.")

}
