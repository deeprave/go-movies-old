package app

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
}

func NewConfig(version string) *Config {
	return &Config{
		Version: version,
		Env:     "prod",
		Port:    9000,
		Host:    "localhost",
	}
}

func NewConfigFromFile(filename string, version string, v ...any) (*Config, error) {
	cfg := NewConfig(version)
	return cfg, cfg.Read(filename, version, v...)
}

func (cfg *Config) VersionOk(version string) bool {
	return version == cfg.Version
}

func (cfg *Config) Read(filename string, version string, v ...any) error {
	filepath, err := FindFile(filename, v...)
	if err == nil {
		var data []byte
		if data, err = os.ReadFile(filepath); err == nil {
			err = yaml.Unmarshal(data, cfg)
		}
		if err == nil && !cfg.VersionOk(version) {
			err = errors.New("config version '%s' is not compatible")
		}
	}
	return err
}
