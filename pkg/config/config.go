package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config Config
type Config struct {
	Driver   string `json:"driver" yaml:"driver"`
	Host     string `json:"host" yaml:"host"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	Prefixes string `json:"prefixes" yaml:"prefixes"`
}

// LoadConfig LoadConfig
func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("./conf/config.json")
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

type MappingItem struct {
	Underlying string `json:"underlying" yaml:"underlying"`
	Nullable   string `json:"nullable" yaml:"nullable"`
}

type MappingConfig = map[string]*MappingItem

var mappingCaches = map[string]MappingConfig{}

// LoadMappingConfig LoadMappingConfig
func LoadMappingConfig(lang string, providerName string) (MappingConfig, error) {
	path := fmt.Sprintf("./conf/language/%s/%s.json", lang, providerName)
	if conf, ok := mappingCaches[path]; ok {
		return conf, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := MappingConfig{}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}

	mappingCaches[path] = conf
	return conf, nil
}
