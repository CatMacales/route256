package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env            string               `yaml:"env" env-default:"prod"`
	Host           string               `yaml:"host" env-default:"localhost"`
	Port           uint32               `yaml:"port" env-default:"8080"`
	ProductService ProductServiceConfig `yaml:"product_service" env-required:"true"`
	LOMSService    LOMSServiceConfig    `yaml:"loms_service" env-required:"true"`
}

type ProductServiceConfig struct {
	URL   string `yaml:"url" env-required:"true"`
	Token string `yaml:"token" env-required:"true"`
}

type LOMSServiceConfig struct {
	URL string `yaml:"url" env-required:"true"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
