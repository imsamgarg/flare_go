package durable

import (
	"encoding/json"
	"flare/internal/models"
	"log"
	"os"
)

var Config models.ConfigInfo

func GetConfig(filepath string) {
	fileContent, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(fileContent, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
