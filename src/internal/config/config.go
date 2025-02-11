package config

import (
	"os"

	"github.com/x-syaifullah-x/go-crud/src/pkg/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type DBConfig struct {
	Host           string                 `yaml:"host"`
	Port           string                 `yaml:"port"`
	User           string                 `yaml:"user"`
	Password       string                 `yaml:"password"`
	Name           string                 `yaml:"name"`
	ConnectionPool DBConnectionPoolConfig `yaml:"connection_pool"`
}

type DBConnectionPoolConfig struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnetcion     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
	MaxIdletimeConnection uint8 `yaml:"max_idletime_connection"`
}

var config *Config

func LoadConfig(filename string) (err error) {
	configBytes, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	return yaml.Unmarshal(configBytes, &config)
}

func GetConfig() Config {
	if config == nil {
		logger.Fatal("Make sure the config file has been loaded.")
	}
	return *config
}
