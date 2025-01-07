package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		fmt.Println(i.Name)
		for _, addr := range addrs {
			fmt.Printf("\t%v\n", addr)
		}
	}
}
