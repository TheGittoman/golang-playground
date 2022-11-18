package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config class
type Config struct {
	Name string   `json:"Name"`
	Type string   `json:"Type"`
	List []string `json:"List"`
}

// ReadConfig output string array from config.json
func ReadConfig() *Config {
	file, err := ioutil.ReadFile("./data/config.json")
	if err != nil {
		log.Println(err)
	}
	config := new(Config)
	err = json.Unmarshal(file, config)
	if err != nil {
		log.Println(err)
	}
	return config
}
