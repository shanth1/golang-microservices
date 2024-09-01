package main

import (
	"fmt"

	"github.com/shanth1/golang-microservices/url-shortener/internal/config"
)

func main() {
	// TODO: init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init jlogger

	// TODO: init storage

	// TODO: init router

	// TODO: run server
}
