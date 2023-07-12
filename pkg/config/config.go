package config

import (
	"encoding/json"
	"os"

	"github.com/temathc/news-aggregator/models"
)

func GetConf() (models.ConfModel, error) {
	confFile, err := os.Open("internal/config/conf.json")
	if err != nil {
		return models.ConfModel{}, err
	}
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	var config models.ConfModel
	err = decoder.Decode(&config)
	if err != nil {
		return models.ConfModel{}, err
	}

	return config, nil
}
