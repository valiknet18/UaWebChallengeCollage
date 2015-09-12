package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ConsumerApiKey string
	ConsumerSecret string
}

func parseConfig() *Config {
	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal("Error open config.json")
	}

	configStruct := &Config{}

	json.Unmarshal(data, configStruct)

	return configStruct
}
