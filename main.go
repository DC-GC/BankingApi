package main

import (
	"fmt"
)

func main() {
	server := NewAPISErver(":3000")
	server.Run()
	fmt.Println("Hello world")
}
