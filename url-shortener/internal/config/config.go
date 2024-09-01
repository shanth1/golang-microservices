package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	httpServer  HttpServer
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:3000"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout string        `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("config path is not found")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("file is not exist: %s\n", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("read config error: %v\n", err)
	}

	return &config
}
