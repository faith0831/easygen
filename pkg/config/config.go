package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Conf 配置信息
type Conf struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// GetConfig 取配置信息
func GetConfig() (*Conf, error) {
	data, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		return nil, err
	}

	conf := &Conf{}
	err = yaml.Unmarshal([]byte(data), conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
