package config

import (
	"log"
	"os"
	"path/filepath"
)

func ConfigPathInit() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}
	configPath := filepath.Join(currentDir, "./configs/config.yaml")
	os.Setenv("CONFIG_PATH", configPath)
}
