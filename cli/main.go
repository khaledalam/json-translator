package main

import (
	jsontranslator "github.com/khaledalam/json-translator"
	"log"
	"os"
)

/*
* Usage:
*
* $ go run main.go test.json en de
 */

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Json file or from_to langs arguments are missing.")
		os.Exit(0)
	}

	// json file that we want to translate
	jsonFile := os.Args[1]
	langFrom := os.Args[2]
	langTo := os.Args[3]

	// Let's first read the `config.json` file
	content, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal("Error when opening json file: ", err)
	}

	jsonStr := jsontranslator.TranslateJson(content, langFrom, langTo)

	if jsonStr == nil {
		log.Fatal("Error: ", jsonStr)
		os.Exit(2)
	}

	err = os.WriteFile("OutputTranslatedJsonFile.json", jsonStr, 0644)
	if err != nil {
		log.Fatal("Error during saving the translated file: ", err)
		os.Exit(3)
	}

	log.Println("Finish!")
}
