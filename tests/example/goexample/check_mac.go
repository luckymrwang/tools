package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		fmt.Println(inter.Name, inter.HardwareAddr.String())
	}
	fmt.Println("get default mac:", getMacAddr())
}

func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}
