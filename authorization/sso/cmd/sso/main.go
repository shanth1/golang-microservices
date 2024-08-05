package main

import (
	"fmt"

	"github.com/shanth1/golang-microservices/authorization/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
