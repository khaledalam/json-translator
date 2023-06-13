package main

import (
	jsontranslator "github.com/khaledalam/json-translator"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
* Usage:
*
* $ go run main.go test.json 5 en de it
 */

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Json file or languages arguments are missing.")
	}

	// Single translation:
	singleTranslationMain()

	// Multiple translation:
	//multipleTranslationMain()
}

func singleTranslationMain() {
	jsonFile := os.Args[1]
	intervalInMillisecond, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Error: wrong interval argument.", err)
	}
	languageFrom := os.Args[3]
	languageTo := os.Args[4]

	jsonStr := jsontranslator.TranslateJsonFile(jsonFile, intervalInMillisecond, languageFrom, languageTo)

	if jsonStr == nil {
		log.Fatal("Error: ", jsonStr)
	}

	err = os.WriteFile("Translated_"+languageFrom+"_"+languageTo+".json", jsonStr, 0644)
	if err != nil {
		log.Fatal("Error during saving the translated file.", err)
	}

	log.Println("Finish!")
}

func multipleTranslationMain() {
	jsonFile := os.Args[1]
	intervalInMillisecond, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Error: wrong interval argument.", err)
	}
	languageFrom := os.Args[3]
	languagesTo := os.Args[4:]

	jsonStr := jsontranslator.TranslateJsonFiles(jsonFile, intervalInMillisecond, languageFrom, languagesTo...)

	if jsonStr == nil {
		log.Fatal("Error: ", jsonStr)
	}

	err = os.WriteFile("Translated_"+languageFrom+"_"+strings.Join(languagesTo, "-")+".json", jsonStr, 0644)
	if err != nil {
		log.Fatal("Error during saving the translated file.", err)
	}

	log.Println("Finish!")
}
