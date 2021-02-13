package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Server struct {
		BindAddress string `yaml:"bindAddress" envconfig:"SERVER__BIND_ADDRESS"`
		Cors        struct {
			AllowOrigins []string `yaml:"allowOrigins" envconfig:"SERVER__CORS__ALLOW_ORIGINS`
		} `yaml:"cors"`
	} `yaml:"server"`
	Database struct {
		Driver  string `yaml:"driver" envconfig:"DATABASE__DRIVER"`
		Address string `yaml:"address" envconfig:"DATABASE__ADDRESS"`
		Debug   bool   `yaml:"debug" envconfig:"DATABASE__DEBUG"`
	} `yaml:"database"`
	Authentication struct {
		AllowedApiKeys []string `yaml:"allowedApiKeys" envconfig:"AUTHENTICATION__ALLOWED_API_KEYS"`
	} `yaml:"authentication"`
}

var Cfg Config

func ReadConfig(configPath, envPrefix string) {
	readFileConfig(&Cfg, configPath)
	readEnv(&Cfg, envPrefix)

	validateConfig()
}

func validateConfig() {
	if len(Cfg.Authentication.AllowedApiKeys) == 0 {
		panic("missing api keys in configuration")
	}
}

func readFileConfig(cfg *Config, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	} else if err != nil {
		log.Fatalf("error while checking config yaml: %v", err)
	}

	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("error while reading config yaml: %v", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}
}

func readEnv(cfg *Config, prefix string) {
	err := envconfig.Process(prefix, cfg)
	if err != nil {
		log.Fatalf("failed to process environment variables for configuration: %v", err)
	}
}
