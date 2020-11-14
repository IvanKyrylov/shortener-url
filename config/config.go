package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"password"`
	} `json:"server"`
	Postgres struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Db       string `json:"db"`
	} `json:"postgres"`
	Optional struct {
		Prefix string `json:"prefix"`
	} `json:"options"`
}

func FromConfigFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
