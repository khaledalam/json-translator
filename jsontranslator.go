// Package jsontranslator
// Simple GoLang script that helps with translate json files using Google Translator.
// Author: Khaled Alam <khaledalam.net@gmail.com>
package jsontranslator

import (
	"encoding/json"
	gt "github.com/bas24/googletranslatefree"
	iso6391 "github.com/emvi/iso-639-1"
	"log"
	"os"
	"time"
)

// TranslateJsonFile allow to translate single json file from language to another language.
// It returns nil or the translated json object string.
// ____
// Params examples:
// jsonFile:				test.json				desc: original input json file.
// intervalInMillisecond:	5						desc: Sleep time between each key translation.
// languageFrom:			en						desc: input ISO 639‑1 code language.
// languageTo:				de						desc: output/target ISO 639‑1 Code language.
func TranslateJsonFile(jsonFile string, intervalInMillisecond int, languageFrom string, languageTo string) []byte {

	// Read json file
	content, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal("Error: when reading json file: ", err)
		return nil
	}

	// Validate languageFrom ISO code
	if iso6391.ValidCode(languageFrom) != true {
		log.Fatal("Error: invalid languageFrom code, use ISO 639‑1 codes ", err)
		return nil
	}

	// Validate languageTo ISO code
	if iso6391.ValidCode(languageTo) != true {
		log.Fatal("Error: invalid languageTo code, use ISO 639‑1 codes ", err)
		return nil
	}

	if intervalInMillisecond < 0 || intervalInMillisecond > 10000 {
		log.Fatal("Error: invalid interval range should be [0 : 10000].", err)
		return nil
	}

	// Map to parse original json file into it per each key
	// inputMap[x] = y
	inputMap := map[string]string{}

	// Map to save translation pairs <key: val>
	// outputMap[x] = z
	outputMap := map[string]string{}

	// Parse json file content
	err = json.Unmarshal(content, &inputMap)
	if err != nil {
		log.Fatal("Error: during parsing json file: ", err)
		return nil
	}

	// Iterate over all json object pairs
	ii := 1
	for key, _ := range inputMap {
		result, _ := gt.Translate(key, languageFrom, languageTo)
		time.Sleep(time.Millisecond * time.Duration(intervalInMillisecond))
		outputMap[key] = result
		log.Println("["+languageFrom+" => "+languageTo+"] Done: ", ii, "/", len(inputMap))
		ii += 1
	}

	// Convert translated map to json object
	jsonResult, err := json.MarshalIndent(outputMap, "", "  ")
	if err != nil {
		log.Fatal("Error: during converting translated map to json", err.Error())
		return nil
	}
	return jsonResult
}

// TranslateJsonFiles allow to translate json file from language to other multiple languages.
// It returns nil or the map of translated json objects strings.
// ____
// Params examples:
// jsonFile:				test.json				desc: original input json file.
// intervalInMillisecond:	5						desc: Sleep time between each key translation.
// languageFrom:			en						desc: input ISO 639‑1 code language.
// languagesTo:				de	it	ar				desc: output/target ISO 639‑1 Code language.
func TranslateJsonFiles(inputJsonFile string, intervalInMillisecond int, languagesFrom string, languagesTo ...string) []byte {

	// Map to save multiple dimension translation pairs <key: val>
	// outputMap[en] = { x: y, ... }
	// outputMap[it] = { x: z, ... }
	outputMap := map[string]map[string]string{}

	for _, languageKey := range languagesTo {

		res := TranslateJsonFile(inputJsonFile, intervalInMillisecond, languagesFrom, languageKey)

		mapVal := map[string]string{}
		err := json.Unmarshal(res, &mapVal)
		if err != nil {
			log.Fatal("Error: during parsing json file [multi] ", err)
			return nil
		}

		outputMap[languageKey] = mapVal
	}

	jsonResult, err := json.MarshalIndent(outputMap, "", "  ")
	if err != nil {
		log.Fatal("Error: during converting translated map to json", err.Error())
		return nil
	}
	return jsonResult
}
