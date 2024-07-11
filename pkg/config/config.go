package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config Config
type Config struct {
	Driver                 string `json:"driver"`
	Host                   string `json:"host"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	Database               string `json:"database"`
	FilteredTablePrefixes  string `json:"filteredTablePrefixes"`
	FilteredCreatedColumns string `json:"filteredCreatedColumns"`
	FilteredUpdatedColumns string `json:"filteredUpdatedColumns"`
}

// LoadConfig LoadConfig
func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("./conf/config.json")
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// SaveConfig SaveConfig
func SaveConfig(c *Config) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("./conf/config.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

type MappingItem struct {
	Underlying string `json:"underlying"`
	Nullable   string `json:"nullable"`
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
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	mappingCaches[path] = conf
	return conf, nil
}
