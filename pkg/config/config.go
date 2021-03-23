package config

import (
	"io/ioutil"

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
	data, err := ioutil.ReadFile("./conf/config.yml")
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
