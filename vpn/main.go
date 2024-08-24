package main

import (
	"fmt"

	"github.com/shanth1/golang-microservices/vpn/vpn"
)

func main() {
	v, err := vpn.NewVpn()
	if err != nil {
		fmt.Printf("NewVpn error: %v\n", err)
		return
	}
	defer v.Close()
	v.Run()

}
