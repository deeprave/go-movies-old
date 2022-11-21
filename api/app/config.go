package app

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Env     string `yaml:"env"`
	Version string `yaml:"version"`
}

func NewConfig(version string) *Config {
	return &Config{
		Port:    9000,
		Host:    "localhost",
		Env:     "prod",
		Version: version,
	}
}

func (cfg *Config) Read(fromFile string) bool {
	filename, err := FindFile(fromFile)
	if err == nil {
		var data []byte
		if data, err = os.ReadFile(filename); err != nil {
			log.Printf("Config '%s': %v", fromFile, err)
		} else {
			err = yaml.Unmarshal(data, cfg)
			if err != nil {
				log.Fatalf("yaml unmarshal error: %v", err)
			}
			return true
		}
	}
	return false
}
