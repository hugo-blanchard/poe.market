package config

import (
	"encoding/json"
	"log"
	"os"
)

func NewConfigFile() ConfigFile {
	jsonFile, error := os.Open("poe.market.config.json")
	if error != nil {
		log.Fatal(error)
	}

	var configFile ConfigFile
	json.NewDecoder(jsonFile).Decode(&configFile)

	return configFile
}

type ConfigFile struct {
	PgConfig PgConfig `json:"postgres"`
	Items    []string `json:"items"`
}

type PgConfig struct {
	DbName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SslMode  string `json:"sslmode"`
}
