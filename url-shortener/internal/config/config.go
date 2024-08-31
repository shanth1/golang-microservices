package config

import "time"

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

func MustLoad() {
	panic("not implemented")
}
