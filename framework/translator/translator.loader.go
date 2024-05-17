package translator

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/yusufocaliskan/tiny-go-mvc/app/utils"
	"github.com/yusufocaliskan/tiny-go-mvc/config"
)

type TranslationEntry struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

type TranslationsMap map[string]TranslationEntry



func LoadErrorTextFile(selectedLang string) TranslationsMap {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("LoadErrorTextFile Error")
	}


	//is the selectedLang one of the acceptaleLangs?
	if(!utils.IsContains(config.AcceptableLangs, selectedLang)){
		selectedLang = config.DefaultLanguage 
	}

	//path
	translationFilePath := filepath.Join(wd, "app/locale/"+selectedLang+".json")

	var translations TranslationsMap
	data, err := os.ReadFile(translationFilePath)
	print("data", data)

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
	if !exists{
		return &TranslationEntry{Text: "Translation not found", Code: ""}
	}
	translations := val.(TranslationsMap)


	if entry, ok := translations[key]; ok {
		return &entry
	}
	return &TranslationEntry{Text: "Translation not found", Code: ""}
}
