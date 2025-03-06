package config

import (
	"fmt"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// ConfigFile is the default config file
var ConfigFile = "./config.yml"

// GlobalConfig is the global config
type Configuration struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// ServerConfig is the server config
type ServerConfig struct {
	Mode               string
	Port               string
	Version            string
	StaticDir          string `yaml:"static_dir"`
	MaxMultipartMemory int64  `yaml:"max_multipart_memory"`
}

// DatabaseConfig is the database config
type DatabaseConfig struct {
	URL          string
	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxOpenConns int `yaml:"max_open_conns"`
}

// global configs
var (
	Global Configuration
)

// Load config from file
func Load(file string) (Configuration, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Printf("%v", err)
		return Global, err
	}

	err = yaml.Unmarshal(data, &Global)
	if err != nil {
		log.Printf("%v", err)
		return Global, err
	}

	return Global, nil
}

// loads configs
func Init() {
	fmt.Println("=================================================")
	fmt.Println("⌛          Loading configurations...          ⌛")
	fmt.Println("=================================================")

	if os.Getenv("config") != "" {
		ConfigFile = os.Getenv("config")
	}

	if _, err := Load(ConfigFile); err != nil {
		log.Fatal("fail to load configs: " + err.Error())
	}
}
