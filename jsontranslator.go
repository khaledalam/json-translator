package jsontranslator

import (
	"encoding/json"
	gt "github.com/bas24/googletranslatefree"
	"log"
	"time"
)

func TranslateJson(content []byte, langFrom string, langTo string) []byte {

	// sleep time between each key translation.
	interval := time.Millisecond * 5

	// map to hold new translation values
	mp := map[string]string{}

	err := json.Unmarshal(content, &mp)
	if err != nil {
		log.Fatal("Error during parsing: ", err)
		return nil
	}

	outputTranslatedJsonFile := map[string]string{}

	ii := 1
	for key, _ := range mp {
		result, _ := gt.Translate(key, langFrom, langTo)
		time.Sleep(interval)
		outputTranslatedJsonFile[key] = result
		log.Println("Done: ", ii, "/", len(mp))
		ii += 1
	}

	// Save translated json object to file
	// @TODO: save translated file keys in same order as original file.
	jsonStr, err := json.MarshalIndent(outputTranslatedJsonFile, "", "  ")
	if err != nil {
		log.Fatal("Error: ", err.Error())
		return nil
	} else {
		return jsonStr
	}

}
