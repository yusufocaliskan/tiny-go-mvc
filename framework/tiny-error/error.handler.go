package tinyerror

import (
	"encoding/json"
	"log"
	"os"
)

type ErrorHandler struct {
	Text string
	Code int
}

func LoadErrorTextFile() interface{} {

	var translations map[string]map[string]string
	data, err := os.ReadFile("../../app/locale/en.json")
	if err != nil {
		log.Fatal("Translation file connot be load")
	}

	err = json.Unmarshal(data, &translations)

	if err != nil {

		log.Fatal("Data connot be binding with translation variable")
	}

	return translations
}
