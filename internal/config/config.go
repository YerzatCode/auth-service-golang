package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port        string `yaml:"port" env-default:":8080"`
	StoragePath string `yaml:"StoragePath" env-required:"true"`
	Secret      string `yaml:"SecretKey"`
}

var Cfg Config

func InitConfig() {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("configpath is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist:%s", configPath)
	}

	if err := cleanenv.ReadConfig(configPath, &Cfg); err != nil {
		log.Fatal("cannot read config:%s", err)
	}

}
