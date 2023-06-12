package main

import (
	jsontranslator "github.com/khaledalam/json-translator"
	"log"
	"os"
	"strings"
)

/*
* Usage:
*
* $ go run main.go test.json en de it
 */

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Json file or languages arguments are missing.")
		os.Exit(0)
	}

	// Single translation:
	//singleTranslationMain()

	// Multiple translation:
	multipleTranslationMain()
}

func singleTranslationMain() {
	jsonFile := os.Args[1]
	languageFrom := os.Args[2]
	languageTo := os.Args[3]

	jsonStr := jsontranslator.TranslateJsonFiles(jsonFile, 5, languageFrom, languageTo)

	if jsonStr == nil {
		log.Fatal("Error: ", jsonStr)
		os.Exit(1)
	}

	err := os.WriteFile("Translated_"+languageFrom+"_"+languageTo+"_"+"JsonFile.json", jsonStr, 0644)
	if err != nil {
		log.Fatal("Error during saving the translated file: ", err)
		os.Exit(2)
	}

	log.Println("Finish!")
}

func multipleTranslationMain() {
	jsonFile := os.Args[1]
	languageFrom := os.Args[2]
	languagesTo := os.Args[3:]

	jsonStr := jsontranslator.TranslateJsonFiles(jsonFile, 5, languageFrom, languagesTo...)

	if jsonStr == nil {
		log.Fatal("Error: ", jsonStr)
		os.Exit(1)
	}

	err := os.WriteFile("Translated_"+languageFrom+"_"+strings.Join(languagesTo, "-")+"_"+"JsonFile.json", jsonStr, 0644)
	if err != nil {
		log.Fatal("Error during saving the translated file: ", err)
		os.Exit(2)
	}

	log.Println("Finish!")

}
