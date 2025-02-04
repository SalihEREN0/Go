package main

import (
	"fmt"
	"net"
)

func NewServer() {
	go fmt.Println("Enter the port number")
	var port string
	go fmt.Scanln(&port)
	Port := ":" + port
	s, err := net.ResolveUDPAddr("udp", Port)
	if err != nil {
		go fmt.Println("Error resolving UDP address:", err)
		return
	}
	conn, err := net.ListenUDP("udp", s)
	if err != nil {
		go fmt.Println("Error setting up UDP connection:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			go fmt.Println("Error reading from UDP connection:", err)
			continue
		}
		go fmt.Printf("Received %s from %s\n", string(buffer[:n]), addr)
	}
}

func main() {
	garb := 0
	fmt.Println("Enter 1 to exit")
	for garb != 1 {
		go NewServer()
		go fmt.Scanln(&garb)
	}
}
