package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

const (
	ConfigPath  = "template/.config"
	SecretsPath = "template/.config.secrets"
)

type Config struct {
	Secrets struct {
		AocSession string `env:"SESSION"`
	} `envPrefix:"AOC_SECRETS_"`

	Public struct {
		CurrentYear string `env:"YEAR"`
	} `envPrefix:"AOC_"`
}

var C Config = Config{}

func init() {
	godotenv.Load(SecretsPath)
	godotenv.Load(ConfigPath)

	if err := env.Parse(&C); err != nil {
		panic(err)
	}
}
