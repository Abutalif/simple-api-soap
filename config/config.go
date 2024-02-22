package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	ApiKey        string `env:"APIKEY"`
	ServerAddress string `env:"SERVER_ADRS"`
	Port          string `env:"PORT"`
}

func GetConfigs(filePath string) (Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(filePath, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
