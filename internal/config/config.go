package config

import (
	"os"

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
	Host           string           `yaml:"host"`
	Port           string           `yaml:"port"`
	User           string           `yaml:"user"`
	Password       string           `yaml:"password"`
	Name           string           `yaml:"name"`
	ConnectionPool DBConnPoolConfig `yaml:"connection_pool"`
}

type DBConnPoolConfig struct {
	MaxIdleConn     uint8 `yaml:"max_idle_connection"`
	MaxOpenConn     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConn uint8 `yaml:"max_lifetime_connection"`
	MaxIdletimeConn uint8 `yaml:"max_idletime_connection"`
}

var Cfg Config

func LoadConfig(filename string) (err error) {
	configByte, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(configByte, &Cfg)
	return
}
