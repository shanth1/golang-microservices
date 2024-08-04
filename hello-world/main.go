package main

import (
	"fmt"

	"github.com/shanth1/golang-microservices/lib1"
)

func main() {
	fmt.Println("Hello2")
	x := lib1.Test()
	fmt.Println(x)
}
