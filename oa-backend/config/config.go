package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var SystemConfig Config

type Config struct {
	Server struct {
		Mode string `yaml:"mode"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Db struct {
		Type     string `yaml:"type"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
}

func init() {
	var configFile []byte
	var err error

	configFile, err = ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("can not find the config file:", err)
	}

	yaml.Unmarshal(configFile, &SystemConfig)
}
