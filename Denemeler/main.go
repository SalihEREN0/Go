package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Enter the port number")
	var port string
	fmt.Scanln(&port)
	Port := ":" + port
	s, err := net.ResolveUDPAddr("udp", Port)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}
	conn, err := net.ListenUDP("udp", s)
	if err != nil {
		fmt.Println("Error setting up UDP connection:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			continue
		}
		fmt.Printf("Received %s from %s\n", string(buffer[:n]), addr)
	}
}
