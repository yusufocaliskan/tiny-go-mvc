package translator

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gptverse/init/app/utils"
	"github.com/gptverse/init/config"
)

type TranslationEntry struct {
	Text string `json:"text"`
	Code string `json:"code"`
}
type TranslationSwaggerResponse struct {
	Message TranslationEntry `json:"message"`
}

type TranslationsMap map[string]TranslationEntry

func LoadErrorTextFile(selectedLang string) TranslationsMap {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("LoadErrorTextFile Error")
	}

	//is the selectedLang one of the acceptaleLangs?
	if !utils.IsContains(config.AcceptableLangs, selectedLang) {
		selectedLang = config.DefaultLanguage
	}

	//path
	translationFilePath := filepath.Join(wd, "app/locale/"+selectedLang+".json")

	var translations TranslationsMap
	data, err := os.ReadFile(translationFilePath)

	if err != nil {
		log.Fatalf("Translation file cannot be loaded: %v", err)
	}

	err = json.Unmarshal(data, &translations)
	if err != nil {
		log.Fatalf("Data cannot be bound with translation variable: %v", err)
	}

	return translations
}

func GetMessage(ctx *gin.Context, key string) *TranslationEntry {

	val, exists := ctx.Get("translations")
	if !exists {
		print("exists", exists)
		return &TranslationEntry{Text: "Translation not found", Code: ""}
	}

	translations := val.(TranslationsMap)

	if entry, ok := translations[key]; ok {

		//Return the actual text
		return &entry
	}

	return &TranslationEntry{Text: key}
}
