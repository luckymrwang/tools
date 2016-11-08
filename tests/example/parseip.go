package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name := "192.168.1.1"
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
